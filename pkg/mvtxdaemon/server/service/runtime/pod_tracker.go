/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runtime

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"

	"github.com/numaproj/numaflow/pkg/apis/numaflow/v1alpha1"
	"github.com/numaproj/numaflow/pkg/shared/logging"
)

// PodTracker tracks the active pods for a MonoVertex.
type PodTracker struct {
	monoVertex          *v1alpha1.MonoVertex
	log                 *zap.SugaredLogger
	httpClient          monitorHttpClient
	activePodsCount     int
	activePodsMutex     sync.RWMutex
	refreshInterval     time.Duration
	firstPodsUpdateChan chan struct{} // Channel to signal the first active pods update is done
}

// NewPodTracker creates a new pod tracker instance.
func NewPodTracker(ctx context.Context, mv *v1alpha1.MonoVertex) *PodTracker {
	pt := &PodTracker{
		monoVertex: mv,
		log:        logging.FromContext(ctx).Named("RuntimePodTracker"),
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second,
		},
		refreshInterval:     30 * time.Second,
		firstPodsUpdateChan: make(chan struct{}),
	}
	return pt
}

// Start starts the pod tracker to track the active pods for the MonoVertex.
func (pt *PodTracker) Start(ctx context.Context) error {
	pt.log.Debugf("Starting tracking active pods for MonoVertex %s...", pt.monoVertex.Name)
	go pt.trackActivePods(ctx)
	return nil
}

func (pt *PodTracker) trackActivePods(ctx context.Context) {
	// start updating active pods as soon as called and then after every refreshInterval
	pt.updateActivePods()
	// close the channel to signal first update
	close(pt.firstPodsUpdateChan)
	ticker := time.NewTicker(pt.refreshInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			pt.log.Infof("Context is cancelled. Stopping tracking active pods for MonoVertex %s...", pt.monoVertex.Name)
			return
		case <-ticker.C:
			pt.updateActivePods()
		}
	}
}

// updateActivePods checks the status of all pods and updates the count of activePods accordingly.
func (pt *PodTracker) updateActivePods() {
	var wg sync.WaitGroup
	// Use atomic operations to safely update the maxActiveIndex across multiple goroutines.
	var maxActiveIndex atomic.Int32
	// Initialize maxActiveIndex to -1 to indicate no active pods.
	maxActiveIndex.Store(int32(-1))
	for i := range int(pt.monoVertex.Spec.Scale.GetMaxReplicas()) {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			podName := fmt.Sprintf("%s-mv-%d", pt.monoVertex.Name, index)
			if pt.isActive(podName) {
				for {
					// Load the current value of maxActiveIndex atomically.
					currentMax := maxActiveIndex.Load()
					// checks if the currentMax is less than index.
					if int32(index) > currentMax {
						// checks if maxActiveIndex still holds the value currentMax.
						// atomically updates maxActiveIndex to the new, higher index and returns true.
						if maxActiveIndex.CompareAndSwap(currentMax, int32(index)) {
							break
						}
					} else {
						break
					}
				}
			}
		}(i)
	}
	wg.Wait()
	// Update the active pods count based on the maxActiveIndex.
	pt.setActivePodsCount(int(maxActiveIndex.Load() + 1))
}

func (pt *PodTracker) isActive(podName string) bool {
	headlessSvc := pt.monoVertex.GetHeadlessServiceName()
	// example for 0th pod: https://simple-mono-vertex-mv-0.simple-mono-vertex-mv-headless.default.svc:2470/runtime/errors
	url := fmt.Sprintf("https://%s.%s.%s.svc:%v/runtime/errors", podName, headlessSvc, pt.monoVertex.Namespace, v1alpha1.MonoVertexMonitorPort)
	resp, err := pt.httpClient.Head(url)
	if err != nil {
		pt.log.Debugf("Sending HEAD request to pod %s is unsuccessful: %v, treating the pod as inactive", podName, err)
		return false
	}
	pt.log.Debugf("Sending HEAD request to pod %s is successful, treating the pod as active", podName)
	_ = resp.Body.Close()
	return true
}

// setActivePodsCount sets the activePodsCount.
func (pt *PodTracker) setActivePodsCount(count int) {
	pt.activePodsMutex.Lock()
	defer pt.activePodsMutex.Unlock()
	pt.log.Debugf("Setting active pods count to %d", count)
	pt.activePodsCount = count
}

// GetActivePodsCount returns the number of active pods.
func (pt *PodTracker) GetActivePodsCount() int {
	pt.activePodsMutex.RLock()
	defer pt.activePodsMutex.RUnlock()
	return pt.activePodsCount
}
