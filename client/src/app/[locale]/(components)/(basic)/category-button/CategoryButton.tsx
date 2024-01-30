import { UnactiveButtonColor } from "@/app/[locale]/constants";
import React, { ReactNode } from "react";

// Constants
import style from "./CategoryButton.styled";

interface CountryButtonProps {
  icon: ReactNode;
  name?: string;
  color?: string;
  active?: boolean;
  onClick?: () => void;
}

const CategoryButton: React.FC<CountryButtonProps> = ({
  icon,
  name,
  color,
  active = false,
  onClick,
}) => {
  return (
    <style.GridMain
      item
      xs={3.5}
      style={{ borderColor: active ? color : UnactiveButtonColor }}
      onClick={onClick}
    >
      {/* Icon */}
      <style.DivIcon
        style={{ backgroundColor: active ? color : UnactiveButtonColor }}
      >
        {icon}
      </style.DivIcon>

      {/* Category Name */}
      <style.IconWord
        fontSize="15px"
        fontWeight={400}
        color={active ? "black" : UnactiveButtonColor}
      >
        {name}
      </style.IconWord>
    </style.GridMain>
  );
};

export default CategoryButton;
