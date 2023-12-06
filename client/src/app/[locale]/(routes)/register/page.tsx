"use client";
import React, { useState } from "react";
import {
  Typography,
  IconButton,
  TextField,
  InputAdornment,
  Input,
  FormControl,
  InputLabel,
  Button,
} from "@mui/material";
import axios from "axios";

// Icons
import ArrowBackIosNewRoundedIcon from "@mui/icons-material/ArrowBackIosNewRounded";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";

// Constants
import { DefaultButtonColor, HomeBackgroundColor } from "../../constants";

interface FormTextFieldProps {
  fieldTitle: string;
  type?: string;
  desc?: string;
  value?: string;
  setFunc?: (val: string) => void;
}

const FormTextField: React.FC<FormTextFieldProps> = ({
  fieldTitle,
  type = "TextField",
  desc,
  value,
  setFunc = () => {},
}) => {
  let inputComponent;
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword((show) => !show);

  const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFunc(event.target.value);
  };

  switch (type) {
    case "TextField":
      inputComponent = (
        <TextField
          id="standard-basic"
          label={fieldTitle}
          variant="standard"
          fullWidth
          value={value}
          onChange={handleInput}
        />
      );
      break;
    case "PasswordField":
      inputComponent = (
        <FormControl fullWidth variant="standard">
          <InputLabel htmlFor="standard-adornment-password">
            {fieldTitle}
          </InputLabel>
          <Input
            id="standard-adornment-password"
            variant="standard"
            type={showPassword ? "text" : "password"}
            endAdornment={
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleClickShowPassword}
                  edge="end"
                >
                  {showPassword ? <VisibilityOff /> : <Visibility />}
                </IconButton>
              </InputAdornment>
            }
            onChange={handleInput}
          />
          <Typography color="gray" fontSize="0.8rem" sx={{ marginTop: "2%" }}>
            {desc}
          </Typography>
        </FormControl>
      );
      break;
  }

  return <div style={{ marginBottom: "20%" }}>{inputComponent}</div>;
};

const RegisterPage = () => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [confirmPassword, setConfirmPassword] = useState<string>("");

  const submitForm = async () => {
    const result = await axios.post("/api/v1/users", {
      Email: "",
      Password: "",
    });
    console.log(result);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        backgroundColor: HomeBackgroundColor,
        height: "100vh",
      }}
    >
      {/* Header */}
      <div
        style={{
          width: "100%",
          height: "5%",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          zIndex: 1,
          backgroundColor: HomeBackgroundColor,
        }}
      >
        <IconButton
          onClick={() => {
            window.location.href = "/login";
          }}
        >
          <ArrowBackIosNewRoundedIcon />
        </IconButton>
        <Typography sx={{ flexGrow: 1, textAlign: "center" }}>
          Create Account
        </Typography>
      </div>

      {/* Form Body */}
      <div
        style={{
          marginTop: "15%",
          display: "flex",
          flexDirection: "column",
          width: "85%",
        }}
      >
        <FormTextField fieldTitle="Email" value={email} setFunc={setEmail} />
        <FormTextField
          fieldTitle="Password"
          type="PasswordField"
          desc="* 8-20 characters including uppercase, 
lowercase, number, and special character"
          value={password}
          setFunc={setPassword}
        />
        <FormTextField
          fieldTitle="Confirm Password"
          type="PasswordField"
          value={confirmPassword}
          setFunc={setConfirmPassword}
        />
      </div>

      {/* Next Button */}
      <div
        style={{
          flexGrow: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "flex-end",
          alignItems: "center",
          width: "90%",
          marginBottom: "20%",
        }}
      >
        <Button
          disabled
          onClick={submitForm}
          style={{
            backgroundColor: DefaultButtonColor,
            color: "white",
            borderRadius: "25px",
            width: "90%",
            padding: "10px 0px",
          }}
        >
          Next
        </Button>
      </div>
    </div>
  );
};

export default RegisterPage;
