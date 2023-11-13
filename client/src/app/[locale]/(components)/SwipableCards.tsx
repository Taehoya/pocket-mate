import React from "react";
import { Box, Card, CardContent, Typography } from "@mui/material";
import SwipeableViews from "react-swipeable-views";

const cards = [
  { title: "Card 1", content: "This is the first card." },
  { title: "Card 2", content: "This is the second card." },
  { title: "Card 3", content: "This is the third card." },
];

const SwipeableCards = () => {
  const [activeIndex, setActiveIndex] = React.useState(0);

  const handleIndexChange = (index: number) => {
    setActiveIndex(index);
  };

  const cardStyle = {
    width: "300px",
    height: "400px",
    borderRadius: "20px",
    overflow: "hidden",
    transition: "transform 0.3s",
    position: "relative",
    backgroundColor: "lightblue",
  };

  const activeCardStyle = {
    ...cardStyle,
    boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
  };

  return (
    <Box>
      <SwipeableViews
        index={activeIndex}
        onChangeIndex={handleIndexChange}
        style={{ padding: "10px 30px" }}
        slideStyle={{
          padding: "20px 0px",
          display: "flex",
          justifyContent: "center",
          overflow: "hidden",
        }}
      >
        {cards.map((card, index) => (
          <div
            key={index}
            style={index === activeIndex ? activeCardStyle : cardStyle}
          >
            <CardContent>
              <Typography variant="h5">{card.title}</Typography>
              <Typography>{card.content}</Typography>
            </CardContent>
          </div>
        ))}
      </SwipeableViews>
    </Box>
  );
};

export default SwipeableCards;
