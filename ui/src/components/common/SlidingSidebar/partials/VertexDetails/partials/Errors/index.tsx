import React, { SyntheticEvent, useCallback, useEffect, useState } from "react";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Autocomplete from "@mui/material/Autocomplete";
import TextField from "@mui/material/TextField";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import { CollapsableError } from "./partials/CollapsableError";
import { ErrorDetails } from "../../../../../../../types/declarations/pods";

import "./style.css";

interface ErrorsProps {
  details?: ErrorDetails[];
  square?: boolean;
}

export const Errors = ({ details, square }: ErrorsProps) => {
  const [filteredDetails, setFilteredDetails] = useState<ErrorDetails[]>([]);
  const [podList, setPodList] = useState<string[]>(["All"]);
  const [containerList, setContainerList] = useState<string[]>(["All"]);
  const [selectedPod, setSelectedPod] = useState<string>("All");
  const [selectedContainer, setSelectedContainer] = useState<string>("All");

  const extractUniqueItems = (items: string[]) => Array.from(new Set(items));

  // TODO: fix extraction logic
  const updateLists = useCallback((details?: ErrorDetails[]) => {
    if (details) {
      const pods = extractUniqueItems(details.map((d) => d.pod));
      const containers = extractUniqueItems(details.map((d) => d.container));
      setPodList(["All", ...pods]);
      setContainerList(["All", ...containers]);
    } else {
      setPodList(["All"]);
      setContainerList(["All"]);
    }
  }, []);

  useEffect(() => {
    updateLists(details);
  }, [details, updateLists]);

  const filterDetails = useCallback(
    (
      selectedPod: string,
      selectedContainer: string,
      details?: ErrorDetails[]
    ) => {
      if (!details) return [];
      else if (selectedPod === "All" && selectedContainer === "All")
        return details;
      else if (selectedContainer === "All") {
        return details.filter((d) => d.pod === selectedPod);
      } else if (selectedPod === "All") {
        return details.filter((d) => d.container === selectedContainer);
      } else {
        return details.filter(
          (d) => d.pod === selectedPod && d.container === selectedContainer
        );
      }
    },
    []
  );

  useEffect(() => {
    const filtered = filterDetails(selectedPod, selectedContainer, details);
    // sort by timestamp
    const sorted = [...filtered].sort(
      (a, b) =>
        new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime()
    );
    setFilteredDetails(sorted);
  }, [selectedPod, selectedContainer, details, filterDetails]);

  const onPodsChange = useCallback(
    (_event: SyntheticEvent, newValue: string) => {
      setSelectedPod(newValue);
    },
    []
  );

  const onContainerChange = useCallback(
    (_event: SyntheticEvent, newValue: string) => {
      setSelectedContainer(newValue);
    },
    []
  );

  if (!details || details.length === 0)
    return (
      <Box className={"vertex-errors-paper-container"}>
        <Paper className={"vertex-errors-paper"} square={square} elevation={0}>
          <Box className={"vertex-no-errors"}>No errors found</Box>
        </Paper>
      </Box>
    );

  return (
    <Box className={"vertex-errors-paper-container"}>
      <Paper className={"vertex-errors-paper"} square={square} elevation={0}>
        {/*dropdown filters*/}
        <Box className={"vertex-errors-selector-dropdown"}>
          <Box>
            <Box className={"vertex-dropdown-title"}>Pod</Box>
            {/*pod dropdown*/}
            <Autocomplete
              options={podList}
              disablePortal
              disableClearable
              id="error-pod-select"
              ListboxProps={{
                sx: { fontSize: "1.6rem" },
              }}
              sx={{
                width: "35rem",
                border: "1px solid #E0E0E0",
                borderRadius: "0.5rem",
                "& .MuiOutlinedInput-root": {
                  borderRadius: "0.5rem",
                },
              }}
              autoHighlight
              onChange={onPodsChange}
              value={selectedPod}
              renderInput={(params) => (
                <TextField
                  {...params}
                  variant="outlined"
                  id="outlined-basic"
                  inputProps={{
                    ...params.inputProps,
                    style: { fontSize: "1.6rem" },
                  }}
                  placeholder={"Select a pod"}
                />
              )}
              popupIcon={<KeyboardArrowDownIcon sx={{ fontSize: "3rem" }} />}
            />
          </Box>
          <Box>
            <Box className={"vertex-dropdown-title"}>Container</Box>
            {/*container dropdown*/}
            <Autocomplete
              options={containerList}
              disablePortal
              disableClearable
              id="error-container-select"
              ListboxProps={{
                sx: { fontSize: "1.6rem" },
              }}
              sx={{
                width: "35rem",
                border: "1px solid #E0E0E0",
                borderRadius: "0.5rem",
                "& .MuiOutlinedInput-root": {
                  borderRadius: "0.5rem",
                },
              }}
              autoHighlight
              onChange={onContainerChange}
              value={selectedContainer}
              renderInput={(params) => (
                <TextField
                  {...params}
                  variant="outlined"
                  id="outlined-basic"
                  inputProps={{
                    ...params.inputProps,
                    style: { fontSize: "1.6rem" },
                  }}
                  placeholder={"Select a container"}
                />
              )}
              popupIcon={<KeyboardArrowDownIcon sx={{ fontSize: "3rem" }} />}
            />
          </Box>
        </Box>

        {filteredDetails.length === 0 && (
          <Box>No errors for the selected filters</Box>
        )}
        {filteredDetails.length > 0 && (
          <>
            <Box className={"vertex-errors-table-title"}>
              <Box sx={{ width: "5rem" }} />
              <Box className={"vertex-error-common-title-text"}>Pod Name</Box>
              <Box className={"vertex-error-common-title-text"}>Container</Box>
              <Box
                className={"vertex-error-common-title-text"}
                sx={{ flexGrow: 1 }}
              >
                Message
              </Box>
              <Box className={"vertex-error-common-title-text"}>
                Last Occurred
              </Box>
            </Box>
            {filteredDetails.map((d, idx) => (
              <Box key={`container-${idx}`} sx={{ my: "0.5rem" }}>
                <CollapsableError detail={d} />
              </Box>
            ))}
          </>
        )}
      </Paper>
    </Box>
  );
};
