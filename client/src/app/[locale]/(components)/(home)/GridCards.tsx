import React from "react";
import Image from 'next/image';
import { Box, Typography } from "@mui/material";
import TripObject from "../(object-types)/TripObject";
import SpringNote from "../(notes)/SpringNote";
import { CSSProperties } from "styled-components";

interface GridCardProps {
  trips: TripObject[] | undefined;
}

const GridCards: React.FC<GridCardProps> = ({ trips }) => {
  const cardStyle: CSSProperties = {
    width: "30%",
    height: "170px",
    borderRadius: "20px",
    margin: "5px 5px",
    position: "relative",
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
          <div>
            <SpringNote />
          </div>
          <div
            style={{
              position: "absolute",
              top: 0,
              margin: "0 5px",
              padding: "0px 20px",
              paddingTop: "10px",
              display: "flex",
              flexDirection: "column",
              height: "100%",
            }}
          >
            <Image
              src={`https://flagsapi.com/${trip?.countryProperty?.code}/flat/64.png`}
              alt={`Flag of ${trip?.countryProperty?.code}`}
              width={32}
              height={32}
            />
            <div
              style={{
                flexGrow: 1,
                alignItems: "flex-end",
              }}
            >
              <Typography fontSize="0.8rem" color="white">
                {trip.name}
              </Typography>
            </div>
          </div>
        </div>
      ))}
    </Box>
  );
};

export default GridCards;
