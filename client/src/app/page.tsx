"use client";

import React, { useState } from "react";
import MultiPageForm from "./(page-components)/MultiPageForm";
import {
  Button,
  Dialog,
  IconButton,
  MenuItem,
  Select,
  Typography,
} from "@mui/material";

// ICONS
import SettingsIcon from "@mui/icons-material/Settings";
import WindowIcon from "@mui/icons-material/Window";
import SwipableCards from "./(components)/SwipableCards";

export default function Home() {
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [dropdownValue, setDropdownValue] = useState("coming-soon");

  const addTravelNote = () => {
    setIsFormOpen(true);
  };

  const closeTravelNote = () => {
    setIsFormOpen(false);
  };

  const handleDropdown = (event: React.ChangeEvent<{ value: unknown }>) => {
    setDropdownValue(event.target.value as string);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100vh",
      }}
    >
      {/* Header */}
      <div
        style={{
          marginTop: "3%",
          textAlign: "center",
        }}
      >
        {/* Title Text */}
        <Typography style={{ fontSize: "1.1rem" }}>Travel note list</Typography>

        {/* Button Row */}
        <div
          style={{
            marginTop: "7%",
            padding: "0px 2%",
            display: "flex",
            justifyContent: "space-between",
          }}
        >
          {/* Settings */}
          <IconButton>
            <SettingsIcon />
          </IconButton>

          {/* Dropdown */}
          <Select
            value={dropdownValue}
            onChange={handleDropdown}
            style={{
              width: "190px",
              height: "40px",
              borderRadius: "25px",
              backgroundColor: "white",
              color: "black",
              boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.2)",
              border: "none",
            }}
          >
            <MenuItem value="coming-soon">Coming Soon</MenuItem>
            <MenuItem value="past">Past</MenuItem>
          </Select>

          {/* Change View */}
          <IconButton>
            <WindowIcon />
          </IconButton>
        </div>
      </div>

      {/* Body */}
      <div style={{marginTop: "10%"}}>
        <SwipableCards />
      </div>

      {/* Add Travel Button */}
      <div
        style={{
          flex: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "flex-end",
          alignItems: "center",
          width: "100%",
          marginBottom: "15%",
        }}
      >
        <Button
          onClick={addTravelNote}
          style={{
            backgroundColor: "#fb3f04",
            color: "white",
            borderRadius: "25px",
            width: "90%",
            padding: "10px 0px",
          }}
        >
          Add a Travel Note
        </Button>
      </div>

      {/* New Note Modal */}
      <Dialog
        open={isFormOpen}
        fullScreen
        onClose={closeTravelNote}
        PaperProps={{
          style: {
            overflowY: "hidden",
            maxHeight: "100%",
            backgroundColor: '#f8f7f2'
          },
        }}
      >
        <MultiPageForm closeForm={closeTravelNote} />
      </Dialog>
    </div>
  );
}