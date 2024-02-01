import React, { useRef, useState } from "react";
import {
  FormControl,
  MenuItem,
  Select,
  SelectChangeEvent,
  TextField,
  Typography,
} from "@mui/material";
import { styled } from "@mui/styles";

interface InputAmountProps {
  amount: number;
}

const NoBorderTextField = styled(TextField)({
  "& .MuiOutlinedInput-input": {
    padding: 0,
    paddingTop: "2%",
    color: "#999",
    fontSize: "32px",
    fontWeight: 400,
  },
  "& .MuiOutlinedInput-root": {
    "& fieldset": {
      border: "none",
      outline: "none",
    },
    "&:hover fieldset": {
      border: "none",
      outline: "none",
    },
    "&.Mui-focused fieldset": {
      border: "none",
      outline: "none",
    },
  },
  "& ::placeholder": {
    color: "#999",
  },
});

const InputAmount: React.FC<InputAmountProps> = ({ amount }) => {
  const [currency, setCurrency] = useState("KRW (Won)");
  const textFieldRef = useRef<HTMLInputElement>(null);

  const handleCurrencyDropdown = (event: SelectChangeEvent) => {
    setCurrency(event.target.value);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        backgroundColor: "#F9F9F9",
        minHeight: "120px",
        width: "100%",
        marginTop: "5%",
      }}
    >
      {/* Inner Margin Modififer */}
      <div style={{ margin: "3% 5%" }}>
        {/* Curreny Toggle */}
        <Select
          value={currency}
          onChange={handleCurrencyDropdown}
          style={{
            height: "25%",
            width: "40%",
            fontWeight: 500,
            fontSize: "14px",
          }}
        >
          <MenuItem value="KRW (Won)">KRW (Won)</MenuItem>
        </Select>

        {/* Amount Input Value */}
        <NoBorderTextField inputRef={textFieldRef} placeholder="Enter amount" />

        {/* Converted Currency Value */}
        <Typography fontSize="12px" color="#CCC">= 0 Won ( 910 won per 100 yen )</Typography>
      </div>
    </div>
  );
};

export default InputAmount;
