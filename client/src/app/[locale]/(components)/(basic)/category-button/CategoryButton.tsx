import React, { ReactNode } from "react";

import style from "./CategoryButton.styled";

interface CountryButtonProps {
  icon: ReactNode;
  name?: string;
  color?: string;
  active?: boolean;
}

const CategoryButton: React.FC<CountryButtonProps> = ({
  icon,
  name,
  color,
  active = false,
}) => {
  return (
    <style.GridMain item xs={3.5}>
      {/* Icon */}
      <style.DivIcon style={{ backgroundColor: color }}>{icon}</style.DivIcon>

      {/* Category Name */}
      <style.IconWord fontSize="15px" fontWeight={400}>
        {name}
      </style.IconWord>
    </style.GridMain>
  );
};

export default CategoryButton;
