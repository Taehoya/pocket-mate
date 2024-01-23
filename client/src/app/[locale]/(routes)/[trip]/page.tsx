"use client";

import { useState } from "react";
import { Typography, IconButton, Grid } from "@mui/material";

// Icons
import ArrowBackIosNewRoundedIcon from "@mui/icons-material/ArrowBackIosNewRounded";
import MenuRoundedIcon from "@mui/icons-material/MenuRounded";

// Constants

// Components
import TripFooter from "../../(components)/(trip)/TripFooter";
import TripPlan from "../../(components)/(trip)/TripPlan";
import TripInfo from "../../(components)/(trip)/TripInfo";

const TripPage = ({ params }: { params: { trip: string } }) => {
  const [activeComponent, setActiveComponent] = useState<number>(2);

  const handleSelect = (componentNumber: number) => {
    setActiveComponent(componentNumber);
  };

  const renderActiveComponent = () => {
    switch (activeComponent) {
      case 1:
        return <TripInfo />;
      case 2:
        return <TripPlan />;
      default:
        return <TripPlan />;
    }
  };

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
          Travel Note
        </Typography>
        <IconButton onClick={() => {}}>
          <MenuRoundedIcon />
        </IconButton>
      </div>

      {/* Body Section */}
      <div
        style={{
          flex: 1,
          overflowY: "auto",
        }}
      >
        {renderActiveComponent()}
      </div>

      {/* Footer Buttons */}
      <TripFooter onSelect={handleSelect} activeComponent={activeComponent} />
    </div>
  );
};

export default TripPage;
