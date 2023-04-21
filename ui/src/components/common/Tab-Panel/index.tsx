import React from "react";
import Box from "@mui/material/Box";
import { TabPanelProps } from "../../../utils/types/declarations/tabpanel";

export default function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`info-tabpanel-${index}`}
      aria-labelledby={`info-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: "1.5rem" }}>{children}</Box>}
    </div>
  );
}
