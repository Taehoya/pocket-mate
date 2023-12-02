import React from "react";
import { Box, CardContent, Typography } from "@mui/material";
import TripObject from "../(object-types)/TripObject";

interface GridCardProps {
  trips: TripObject[] | undefined;
}

const GridCards: React.FC<GridCardProps> = ({ trips }) => {
  const cardStyle = {
    width: "30%",
    height: "170px",
    borderRadius: "20px",
    backgroundColor: "lightblue",
    margin: "5px 5px",
  };

  return (
    <Box
      style={{
        width: "100%",
        maxHeight: "100%",
        display: "flex",
        flexWrap: "wrap",
        alignItems: "flex-start",
        overflowY: "scroll",
      }}
    >
      {trips?.map((trip: TripObject, index: number) => (
        <div key={index} style={cardStyle}>
          <CardContent style={{ fontSize: "1rem" }}>
            <Typography>{trip.title}</Typography>
            <Typography>{trip.description}</Typography>
          </CardContent>
        </div>
      ))}
    </Box>
  );
};

export default GridCards;
