import React, { useState } from "react";
import { Box, Typography } from "@mui/material";
import Avatar from "@mui/material/Avatar";
import AvatarGroup from "@mui/material/AvatarGroup";
import SwipeableViews from "react-swipeable-views";
import TripObject from "../(object-types)/TripObject";

// CONSTANTS
import { HomeBackgroundColor, DefaultButtonColor } from "../constants";
import GroupAddIcon from "@mui/icons-material/GroupAdd";

interface SwipeableCardsProps {
  trips: TripObject[] | undefined;
}

const SwipeableCards: React.FC<SwipeableCardsProps> = ({ trips }) => {
  const [activeIndex, setActiveIndex] = useState(0);

  const handleIndexChange = (index: number) => {
    setActiveIndex(index);
  };

  const cardStyle = {
    display: "flex",
    justifyContent: "flex-end", // to make the content be on right side
    width: "92%",
    height: "400px",
    borderRadius: "20px",
    overflow: "hidden",
    transition: "transform 0.3s",
    backgroundRepeat: "no-repeat",
    backgroundPosition: "center",
    backgroundColor: HomeBackgroundColor,
    color: "#E5E5E5",
  };

  return (
    <Box style={{ height: "100%" }}>
      <SwipeableViews
        index={activeIndex}
        onChangeIndex={handleIndexChange}
        style={{ padding: "10px 32px", height: "100%" }}
        slideStyle={{
          padding: "20px 0px",
          display: "flex",
          justifyContent: "center",
          overflow: "hidden",
          height: "100%",
        }}
      >
        {trips?.map((trip: TripObject, index: number) => {
          const startDate = new Date(trip.startDateTime);
          const endDate = new Date(trip.endDateTime);
          const duration = 1 + (endDate.getTime() - startDate.getTime()) / 24;
          return (
            <div
              key={index}
              style={{
                height: "100%",
                width: "100%",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}
            >
              {/* Trip Book */}
              <div
                style={{
                  ...cardStyle,
                  backgroundImage: `url('/trip/SpringNote.svg')`,
                }}
              >
                {/* Section withou Binder */}
                <div
                  style={{
                    width: "85%",
                    height: "100%",
                    paddingRight: "10%",
                    paddingLeft: "20%",
                    display: "flex",
                    flexDirection: "column",
                  }}
                >
                  {/* Head Section */}
                  <div
                    style={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "space-between",
                      marginTop: "20%",
                    }}
                  >
                    <img src={`https://flagsapi.com/${trip?.countryProperty?.code}/flat/64.png`} />
                    <Typography>{duration} days</Typography>
                  </div>

                  {/* Body Section */}
                  <Typography
                    variant="h6"
                    color="white"
                    sx={{
                      flex: 4,
                      display: "flex",
                      alignItems: "center",
                    }}
                  >
                    {trip.name}
                  </Typography>

                  {/* Cash Section */}
                  <Typography
                    variant="h4"
                    color="white"
                    sx={{
                      flex: 1,
                    }}
                  >
                    $ {trip.budget}
                  </Typography>

                  {/* Footer Section */}
                  <div
                    style={{
                      flex: 1,
                    }}
                  >
                    <Typography fontSize="0.8rem" color="white">
                      {startDate.toLocaleDateString()} -{" "}
                      {endDate.toLocaleDateString()}
                    </Typography>
                  </div>
                </div>
              </div>

              {/* Group */}
              <div
                style={{
                  width: "80%",
                  height: "30px",
                  marginTop: "10%",
                  display: "flex",
                  justifyContent: "center",
                }}
              >
                <AvatarGroup total={8}>
                  <Avatar alt="Remy Sharp" src="/" sx={{ bgcolor: "green" }} />
                  <Avatar
                    alt="Agnes Walker"
                    src="/"
                    sx={{ bgcolor: "lightblue" }}
                  />
                  <Avatar
                    alt="Trevor Henderson"
                    src="/"
                    sx={{ bgcolor: "red" }}
                  />
                </AvatarGroup>
                <Avatar
                  sx={{ marginLeft: "3%", backgroundColor: DefaultButtonColor }}
                >
                  <GroupAddIcon />
                </Avatar>
              </div>
            </div>
          );
        })}
      </SwipeableViews>
    </Box>
  );
};

export default SwipeableCards;
