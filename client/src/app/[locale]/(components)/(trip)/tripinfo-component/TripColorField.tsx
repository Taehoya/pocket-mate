import { useState } from "react";
import { IconButton, Typography } from "@mui/material";

// ICONS
import ModeEditIcon from "@mui/icons-material/ModeEdit";

const TripColorField = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      {/* Color 1 */}
      <div
        style={{
          display: "flex",
          margin: "2% 0",
        }}
      >
        <Typography fontSize="1.5rem" fontWeight="bold" style={{ flex: 1 }}>
          Binder
        </Typography>
        <div
          style={{
            flex: 1,
            backgroundColor: "blue",
            width: "30%",
            borderRadius: "5px",
          }}
        />
        <div style={{ flex: 1, marginLeft: "5%" }}>
          <IconButton>
            <ModeEditIcon />
          </IconButton>
        </div>
      </div>
      {/* Color 2 */}
      <div style={{ display: "flex" }}>
        <Typography fontSize="1.5rem" fontWeight="bold" style={{ flex: 1 }}>
          Body
        </Typography>
        <div
          style={{
            flex: 1,
            backgroundColor: "blue",
            width: "30%",
            borderRadius: "5px",
          }}
        />
        <div style={{ flex: 1, marginLeft: "5%" }}>
          <IconButton>
            <ModeEditIcon />
          </IconButton>
        </div>
      </div>
    </div>
  );
};

export default TripColorField;
