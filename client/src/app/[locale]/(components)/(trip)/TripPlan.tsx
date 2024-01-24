import { Typography } from "@mui/material";
import CountryMap from "../(extra)/CountryMap";
import TripContent from "./TripContent";

const TripPlan = () => {
  return (
    <div style={{ display: "flex", flexDirection: "column", height: "100%" }}>
      {/* Title Section */}
      <div
        style={{
          flex: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          paddingLeft: "5%",
        }}
      >
        {/* Title */}
        <Typography fontSize="1.8rem">Trip to London</Typography>

        {/* Duration */}
        <Typography fontSize="0.9rem">2023/10/12 - 2023/10/15</Typography>
      </div>

      {/* Map Preview */}
      <div style={{ flex: 2 }}>
        <CountryMap />
      </div>

      {/* Dynamic Main Body */}
      <div style={{ flex: 5 }}>
        <TripContent />
      </div>
    </div>
  );
};

export default TripPlan;
