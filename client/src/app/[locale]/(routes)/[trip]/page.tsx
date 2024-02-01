"use client";

import { useState } from "react";
import { Typography, IconButton, Grid } from "@mui/material";

// Icons
import ArrowBackIosNewRoundedIcon from "@mui/icons-material/ArrowBackIosNewRounded";

// Components
import TripFooter from "../../(components)/(trip)/TripFooter";
import TripPlan from "../../(components)/(trip)/TripPlan";
import TripInfo from "../../(components)/(trip)/TripInfo";
import TransactionTemplate from "../../(components)/(transaction)/transaction-page/TransactionTemplate";
import TripGraph from "../../(components)/(trip)/TripGraph";
import WebWrapper from "../../(wrapper)/WebWrapper";
import TripDrawer from "../../(components)/(trip)/TripDrawer";

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
      case 3:
        return <TransactionTemplate />;
      case 4:
        return <TripGraph />;
      default:
        return <TripPlan />;
    }
  };

  return (
    <WebWrapper>
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

          {/* Drawer */}
          <TripDrawer />
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
    </WebWrapper>
  );
};

export default TripPage;
