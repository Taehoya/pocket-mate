import React, { useState } from "react";
import { IconButton, Typography } from "@mui/material";

// ICONS
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import ColorPickerDialog from "../../(basic)/dialog-color-picker/ColorPickerDialog";

interface TripColorFieldProps {
  active?: boolean;
}

const TripColorField: React.FC<TripColorFieldProps> = ({ active = false }) => {
  const [toggleDialog, setToggleDialog] = useState(false);

  const handleDialogOpen = () => {
    setToggleDialog(true);
  };

  const handleDialogClose = () => {
    setToggleDialog(false);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      {/* Binder Color */}
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
          {active && (
            <IconButton onClick={handleDialogOpen}>
              <ModeEditIcon />
            </IconButton>
          )}
        </div>
      </div>

      {/* Body Color */}
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
          {active && (
            <IconButton onClick={handleDialogOpen}>
              <ModeEditIcon />
            </IconButton>
          )}
        </div>
      </div>

      {/* Color Picker Dialog */}
      <ColorPickerDialog open={toggleDialog} onClose={handleDialogClose} />
    </div>
  );
};

export default TripColorField;
