"use client";

import { Typography, IconButton } from "@mui/material";

// Icons
import ArrowBackIosNewRoundedIcon from "@mui/icons-material/ArrowBackIosNewRounded";
import SettingsIcon from "@mui/icons-material/Settings";

// Constants
import { DefaultButtonColor, UnactiveColor } from "../../constants";

const TripPage = ({ params }: { params: { trip: string } }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        width: "100%",
        height: "100vh",
        backgroundColor: "white",
      }}
    >
      {/* Top Header */}
      <div
        style={{
          flex: 1,
          backgroundColor: "green",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <IconButton
          onClick={() => {
            window.location.href = "/";
          }}
        >
          <ArrowBackIosNewRoundedIcon />
        </IconButton>
        <Typography sx={{ flexGrow: 1, textAlign: "center" }}>
          Trip Title
        </Typography>
      </div>

      {/* Title Section */}
      <div
        style={{
          flex: 2,
          backgroundColor: "red",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        {/* Settings */}
        <IconButton>
          <SettingsIcon />
        </IconButton>

        {/* Title */}
        <Typography fontSize="1.4rem">나만의 힐링 여행</Typography>

        {/* Duration */}
        <Typography fontSize="1rem">2023/10/12-2023/10/15</Typography>
      </div>

      {/* Dynamic Main Body */}
      <div style={{ flex: 12 }}>wut!</div>

      {/* Footer Buttons */}
      <div
        style={{
          flex: 4,
          display: "flex",
          flexDirection: "column",
          backgroundColor: DefaultButtonColor,
          borderRadius: "20px 20px 0 0",
        }}
      >
        {/* Top Buttons */}
        <div
          style={{
            flex: 1,
          }}
        ></div>
        {/* Bottom Buttons */}
        <div
          style={{
            flex: 1.4,
            backgroundColor: UnactiveColor,
            borderRadius: "20px 20px 0 0",
          }}
        ></div>
      </div>
    </div>
  );
};

export default TripPage;
