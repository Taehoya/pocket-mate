import React, { useState } from "react";
import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent,
} from "@mui/material";

// Constants
import { UnactiveColor } from "@/app/[locale]/constants";

interface DropdownBubbleProps {
  label: string;
  content: string[];
}

const DropdownBubble: React.FC<DropdownBubbleProps> = ({ label, content }) => {
  const [dropdownValue, setDropdownValue] = useState(content[0]);

  const handleDropdown = (event: SelectChangeEvent) => {
    setDropdownValue(event.target.value);
  };

  return (
    <FormControl
      style={{
        borderRadius: "10px",
        border: `1px solid ${UnactiveColor}`,
        height: "100%",
        width: "100%",
        position: "relative",
      }}
    >
      <InputLabel
        variant="standard"
        htmlFor="date-select"
        style={{
          fontSize: "18px",
          fontWeight: 300,
          position: "absolute",
          top: "10%",
          left: "5%",
        }}
      >
        {label}
      </InputLabel>
      <Select
        value={dropdownValue}
        onChange={handleDropdown}
        style={{
          width: "100%",
          height: "100%",
          fontWeight: 500,
          fontSize: "18px",
          paddingTop: "20px",
        }}
      >
        {content.map((value, index) => (
          <MenuItem key={index} value={value}>
            {value}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

export default DropdownBubble;
