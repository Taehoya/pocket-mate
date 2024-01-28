"use client";

import React, { ReactNode } from "react";
import { Grid, useMediaQuery } from "@mui/material";
// import { DeviceFrameset } from "react-device-frameset";
import "react-device-frameset/styles/marvel-devices.min.css";

interface WebWrapperProps {
  children: ReactNode;
}

const WebWrapper: React.FC<WebWrapperProps> = ({ children }) => {
  // const isSmallScreen = useMediaQuery("(max-width:600px)"); // Define your breakpoint

  return (
    <Grid container>
      <Grid item xs={0} md={4}></Grid>
      <Grid item xs={12} md={4}>
        {children}
      </Grid>
      <Grid item xs={0} md={4}></Grid>
    </Grid>
  );
};

export default WebWrapper;

{
  /* <DeviceFrameset device="Galaxy Note 8" color="gold" zoom={1}>
{children}
</DeviceFrameset> */
}
