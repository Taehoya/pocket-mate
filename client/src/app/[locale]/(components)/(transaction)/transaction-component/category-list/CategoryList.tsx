import React, { useState } from "react";

import { Grid, Typography } from "@mui/material";
import CategoryButton from "../../../(basic)/category-button/CategoryButton";

// Icons
import HotelIcon from "@mui/icons-material/LocalHotelOutlined";
import CarIcon from "@mui/icons-material/DirectionsCarFilledOutlined";
import FoodIcon from "@mui/icons-material/FastfoodOutlined";
import ShopIcon from "@mui/icons-material/ShoppingCartOutlined";

const CategoryList = () => {
  const [activeIndex, setActiveIndex] = useState<number | null>(0);

  const IconDict: { [key: string]: JSX.Element } = {
    HotelIcon: <HotelIcon />,
    CarIcon: <CarIcon />,
    FoodIcon: <FoodIcon />,
    ShopIcon: <ShopIcon />,
  };

  const handleButtonClick = (index: number) => {
    setActiveIndex(index);
  };

  return (
    // outer wrapping
    <div
      style={{
        flexGrow: 1,
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        margin: "10% 0px",
      }}
    >
      {/* Section Name */}
      <Typography fontSize="20px" style={{ width: "89%", marginBottom: "5%" }}>
        Category
      </Typography>

      {/* Inner Wrapping */}
      <Grid
        container
        rowGap={2}
        columnGap={2}
        style={{
          width: "80%",
        }}
      >
        {example.map((item, index) => (
          <CategoryButton
            key={index}
            color={item.color}
            name={item.name}
            icon={IconDict[item.icon]}
            active={index === activeIndex}
            onClick={() => handleButtonClick(index)}
          />
        ))}
      </Grid>
    </div>
  );
};

export default CategoryList;

const example = [
  {
    color: "#FB3F04",
    icon: "HotelIcon",
    name: "Hotel",
  },
  {
    color: "#2ecc71",
    icon: "CarIcon",
    name: "Rent",
  },
  {
    color: "#008080",
    icon: "FoodIcon",
    name: "Food",
  },
  {
    color: "#e67e22",
    icon: "ShopIcon",
    name: "Shop",
  },
  {
    color: "#8e44ad",
    icon: "CarIcon",
    name: "Activity",
  },
  {
    color: "#3498db",
    icon: "FoodIcon",
    name: "Entertainment",
  },
  {
    color: "#e74c3c",
    icon: "HotelIcon",
    name: "Food",
  },
];
