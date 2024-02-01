import React, { useState } from "react";
import {
  Avatar,
  Button,
  IconButton,
  SwipeableDrawer,
  Typography,
} from "@mui/material";

// Icons
import MenuRoundedIcon from "@mui/icons-material/MenuRounded";
import CloseIcon from '@mui/icons-material/CloseRounded';

const TripDrawer = () => {
  const [open, setOpen] = useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <IconButton onClick={handleDrawerOpen}>
        <MenuRoundedIcon />
      </IconButton>

      {/* Drawer component */}
      <SwipeableDrawer
        anchor="right"
        open={open}
        onOpen={handleDrawerOpen}
        onClose={handleDrawerClose}
        sx={{
          width: 200,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: 200,
            boxSizing: "border-box",
          },
        }}
      >
        <div>
          {/* Drawer Header */}
      <IconButton onClick={handleDrawerClose}>
        <CloseIcon />
      </IconButton>
          <div
            style={{
              width: "100%",
              height: "50px",
              display: "flex",
              alignItems: "center",
              marginTop: "10px"
            }}
          >
            <div style={{ flex: 2, marginLeft: "20px" }}>
              <Typography fontSize="18px">Antonio Branzerto</Typography>
            </div>
            <div style={{ flex: 1 }}>
              <Avatar alt="Guest123" />
            </div>
          </div>

          {/* List of Group Members */}
        </div>
      </SwipeableDrawer>
    </div>
  );
};

export default TripDrawer;
