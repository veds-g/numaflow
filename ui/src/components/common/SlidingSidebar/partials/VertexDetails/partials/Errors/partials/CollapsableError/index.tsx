import React, { useState } from "react";
import moment from "moment";
import Box from "@mui/material/Box";
import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Divider from "@mui/material/Divider";
import { ago } from "../../../../../../../../../utils";
import { ContainerError } from "../../../../../../../../../types/declarations/pods";

import "./style.css";

interface CollapsableErrorProps {
  detail: ContainerError & { pod: string };
}

// removes initial escape sequences from details
const cleanText = (text: string) => {
  const initialEscapeSequences = ["\b\r\u00123", "\b\r\u0012H", "\b\r\u00128"];
  const regex = new RegExp(`^(${initialEscapeSequences.join("|")})`);
  return text.replace(regex, "");
};

const highlightFilePaths = (rawText: string) => {
  const text = cleanText(rawText);
  const filePathRegex = /((?:\/[^\s]+)+\.[a-zA-Z0-9:]+)|(\bat\s+[^\n]+)/g;
  const exclusionList = ["/google.rpc.DebugInfo", "/debug.Stack"];

  return text.split(filePathRegex).map((part, index) => {
    if (filePathRegex.test(part) && !exclusionList.includes(part)) {
      return (
        <span key={index} style={{ color: "blue", fontWeight: "bold" }}>
          {part}
        </span>
      );
    }
    return part;
  });
};

export const CollapsableError = ({ detail }: CollapsableErrorProps) => {
  const [expanded, setExpanded] = useState(false);

  const handleChange =
    (panel: boolean) => (_event: React.SyntheticEvent, isExpanded: boolean) => {
      setExpanded(isExpanded ? panel : false);
    };

  return (
    <Accordion
      className={"collapsable-error-accordion"}
      expanded={expanded}
      onChange={handleChange(true)}
    >
      <AccordionSummary
        expandIcon={<ExpandMoreIcon />}
        aria-controls="container-dropdown-content"
        id="container-dropdown-header"
        sx={{
          height: "7rem",
          "& .MuiAccordionSummary-content": {
            display: "flex",
            alignItems: "center",
            gap: "0.5rem",
            overflow: "scroll",
          },
          "& .MuiAccordionSummary-expandIconWrapper": {
            order: -1,
          },
        }}
      >
        <Box className={"collapsable-error-title-box"}>
          <Box
            className={"collapsable-error-common-title-text"}
            sx={{ ml: "4.6rem" }}
          >
            {detail.pod}
          </Box>
          <Box className={"collapsable-error-common-title-text"}>
            {detail.container}
          </Box>
          <Box
            className={"collapsable-error-common-title-text"}
            sx={{ flexGrow: 1 }}
          >
            {detail.message}
          </Box>
          <Box className={"collapsable-error-common-title-text"}>
            <Box>{ago(new Date(detail.timestamp))}</Box>
            <Box>
              {moment(new Date(detail.timestamp)).calendar(null, {
                sameDay: "[Today at] LT",
                lastDay: "[Yesterday at] LT",
                lastWeek: "[Last] dddd [at] LT",
                sameElse: "MM/DD/YYYY [at] LT",
              })}
            </Box>
          </Box>
        </Box>
      </AccordionSummary>
      <AccordionDetails className={"collapsable-error-accordion-details"}>
        <Box className={"collapsable-error-accordion-details-box"}>
          <Box className={"collapsable-error-accordion-details-title"}>
            Details
          </Box>
          <Divider orientation="vertical" flexItem color={"#878789"} />
          <Box className={"collapsable-error-accordion-details-title-content"}>
            <pre style={{ whiteSpace: "pre-wrap", wordWrap: "break-word" }}>
              {highlightFilePaths(detail.details)}
            </pre>
          </Box>
          <Divider
            orientation="vertical"
            flexItem
            color={"#878789"}
            sx={{ mr: "25rem" }}
          />
        </Box>
      </AccordionDetails>
    </Accordion>
  );
};
