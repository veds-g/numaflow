import React, { useCallback, useContext } from "react";
import Box from "@mui/material/Box";
import { useLocation } from "react-router-dom";
import { SidebarType } from "../../../../common/SlidingSidebar";
import { AppContextProps } from "../../../../../types/declarations/app";
import { AppContext } from "../../../../../App";
import { ViewType } from "../../../../common/SpecEditor";

import "./style.css";

export interface MonoVertexSummaryProps {
  pipelineId: any;
  pipeline: any;
  refresh: () => void;
}

export function MonoVertexSummaryStatus({
  pipelineId,
  pipeline,
  refresh,
}: MonoVertexSummaryProps) {
  const location = useLocation();
  const query = new URLSearchParams(location.search);
  const namespaceId = query.get("namespace") || "";
  const { setSidebarProps } = useContext<AppContextProps>(AppContext);

  const handleUpdateComplete = useCallback(() => {
    refresh();
    if (!setSidebarProps) {
      return;
    }
    // Close sidebar
    setSidebarProps(undefined);
  }, [setSidebarProps, refresh]);

  const handleSpecClick = useCallback(() => {
    if (!namespaceId || !setSidebarProps) {
      return;
    }
    setSidebarProps({
      type: SidebarType.PIPELINE_UPDATE,
      specEditorProps: {
        titleOverride: `View Pipeline: ${pipelineId}`,
        initialYaml: pipeline,
        namespaceId,
        pipelineId,
        viewType: ViewType.READ_ONLY,
        // viewType: isReadOnly ? ViewType.READ_ONLY : ViewType.TOGGLE_EDIT,
        onUpdateComplete: handleUpdateComplete,
      },
    });
  }, [
    namespaceId,
    pipelineId,
    setSidebarProps,
    pipeline,
    handleUpdateComplete,
  ]);

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        marginTop: "0.6rem",
        flexGrow: 1,
        paddingLeft: "1.6rem",
      }}
    >
      <Box sx={{ width: "fit-content" }}>
        <span className="pipeline-status-title">SUMMARY</span>
        <Box
          sx={{ display: "flex", flexDirection: "row", marginTop: "0.5rem" }}
        >
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              marginRight: "1.6rem",
            }}
          >
            <div className="pipeline-summary-text">
              <span className="pipeline-summary-subtitle">Created On: </span>
            </div>
            <div className="pipeline-summary-text">
              <span className="pipeline-summary-subtitle">
                Last Updated On:{" "}
              </span>
            </div>
            {/*<div className="pipeline-summary-text">*/}
            {/*  <span className="pipeline-summary-subtitle">Last Refresh: </span>*/}
            {/*  2023-12-07T02:02:00Z*/}
            {/*</div>*/}
          </Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              marginRight: "6.4rem",
            }}
          >
            <div className="pipeline-summary-text">
              <span>{pipeline?.metadata?.creationTimestamp}</span>
            </div>
            <div className="pipeline-summary-text">
              <span>{pipeline?.status?.lastUpdated}</span>
            </div>
            {/*<div className="pipeline-summary-text">*/}
            {/*  2023-12-07T02:02:00Z*/}
            {/*</div>*/}
          </Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              width: "19.2rem",
            }}
          >
            <div className="pipeline-summary-text">
              <span className="pipeline-summary-subtitle">
                <div
                  className="pipeline-onclick-events"
                  onClick={handleSpecClick}
                  data-testid="pipeline-spec-click"
                >
                  {`View Specs`}
                </div>
              </span>
            </div>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}
