"use client";

import { Box, Button } from "@mui/material";

const TripCreation = () => {
  return (
    <div
      style={{
        width: "100%",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Box sx={{ fontSize: "1rem", marginTop: "20px" }}>Add a travel note</Box>

      <Box
        sx={{
          marginTop: "25%",
          marginLeft: "15%",
          display: "flex",
          alignItems: "flex-start",
          flexDirection: "column",
          width: "100%",
        }}
      >
        <Box fontSize={20}>Welcome to WithTrip! </Box>
        <Box sx={{ fontSize: "1.6rem", fontWeight: "bold" }}>
          Let's make a travel note
        </Box>
      </Box>

      <Button
        style={{
            marginTop: "20%",
          backgroundColor: "#fb3f04",
          color: "white",
          borderRadius: "25px",
          width: "80%",
          padding: "10px 0px"
        }}
      >
        Sounds great
      </Button>
    </div>
  );
};

export default TripCreation;
