import React, { useCallback, useMemo, useState } from "react";
import Tooltip from "@mui/material/Tooltip";
import Box from "@mui/material/Box";
import { MetricsModal } from "./partials/MetricsModal";
import { useMetricsDiscoveryDataFetch } from "../../../utils/fetchWrappers/metricsDiscoveryDataFetch";
import { dimensionReverseMap } from "../../pages/Pipeline/partials/Graph/partials/NodeInfo/partials/Pods/partials/PodDetails/partials/Metrics/utils/constants";

import "./style.css";

interface MetricsModalWrapperProps {
  disableMetricsCharts: boolean;
  namespaceId: string;
  pipelineId: string;
  vertexId: string;
  type: string;
  metricDisplayName: string;
  value: any;
  presets?: any;
}

export function MetricsModalWrapper({
  disableMetricsCharts,
  namespaceId,
  pipelineId,
  vertexId,
  type,
  metricDisplayName,
  value,
  presets,
}: MetricsModalWrapperProps) {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleOpenModal = useCallback(() => {
    setIsModalOpen(true);
  }, []);
  const handleCloseModal = useCallback(() => {
    setIsModalOpen(false);
  }, []);

  const {
    metricsDiscoveryData: discoveredMetrics,
    error: discoveredMetricsError,
    loading: discoveredMetricsLoading,
  } = useMetricsDiscoveryDataFetch({
    objectType: dimensionReverseMap[type],
  });

  const isClickable = useMemo(() => {
    return (
      !discoveredMetricsError &&
      !discoveredMetricsLoading &&
      !disableMetricsCharts
    );
  }, [discoveredMetricsError, discoveredMetricsLoading, disableMetricsCharts]);

  return (
    <Box>
      {isClickable ? (
        <Tooltip
          title={
            <Box sx={{ fontSize: "1rem" }}>
              Click to get more information about the trend
            </Box>
          }
          placement="top-start"
          arrow
        >
          <Box className={"metrics-hyperlink"} onClick={handleOpenModal}>
            {value}
          </Box>
        </Tooltip>
      ) : (
        <Box style={{ width: "fit-content" }}>{value}</Box>
      )}
      <MetricsModal
        isModalOpen={isModalOpen}
        handleCloseModal={handleCloseModal}
        metricDisplayName={metricDisplayName}
        discoveredMetrics={discoveredMetrics}
        namespaceId={namespaceId}
        pipelineId={pipelineId}
        vertexId={vertexId}
        type={type}
        presets={presets}
      />
    </Box>
  );
}
