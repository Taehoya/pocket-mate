import React from "react";

import { Grid, Typography } from "@mui/material";
import CategoryButton from "../../../(basic)/category-button/CategoryButton";

// Icons
import HotelIcon from "@mui/icons-material/LocalHotelOutlined";
import CarIcon from '@mui/icons-material/DirectionsCarFilledOutlined';
import FoodIcon from '@mui/icons-material/FastfoodOutlined';
import ShopIcon from '@mui/icons-material/ShoppingCartOutlined';

const CategoryList = () => {
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
      <Typography fontSize="20px" style={{ width: "85%", marginBottom: "5%" }}>
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
        <CategoryButton icon={<HotelIcon />} name="Hotel" />
        <CategoryButton color="#2ecc71" icon={<CarIcon />} name="Rent" />
        <CategoryButton color="#008080" icon={<FoodIcon />} name="Food" />
        <CategoryButton color="#e67e22" icon={<ShopIcon />} name="Shop" />
        <CategoryButton color="#8e44ad" icon={<CarIcon />} name="Activity" />
        <CategoryButton color="#3498db" icon={<FoodIcon />} name="Entertainment" />
        <CategoryButton color="#e74c3c" icon={<HotelIcon />} name="Food" />
      </Grid>
    </div>
  );
};

export default CategoryList;
