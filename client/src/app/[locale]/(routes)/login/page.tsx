"use client";

import React, { ChangeEvent, useState } from "react";
import Image from "next/image";
import axios from "axios";
import { useTranslations } from "next-intl";
import { Divider, Typography, Button, Link, TextField } from "@mui/material";

// Constants
import { DefaultButtonColor, LoginLinkColor } from "@/app/[locale]/constants";

interface SSOButtonProps {
  image?: string;
  text?: string;
  type?: string;
  backgroundColor: string;
  logic?: Function;
}

const SSOButton: React.FC<SSOButtonProps> = ({
  image,
  backgroundColor,
  text,
  type = "logo",
  logic,
}) => {
  const imageSrc = `/sso/${image}.svg`;
  var button;
  const buttonStyle = {
    backgroundColor: backgroundColor,
    color: "black",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    alignText: "center",
    margin: "0px 2%",
  };
  switch (type) {
    case "logo":
      button = (
        <Button
          variant="contained"
          style={{
            ...buttonStyle,
            borderRadius: "50%",
            width: "50px",
            height: "50px",
            padding: 0,
            overflow: "hidden",
          }}
        >
          <img
            src={imageSrc}
            alt="SSO Icon"
            style={{
              width: "60%",
              height: "60%",
              borderRadius: "50%",
            }}
          />
        </Button>
      );
      break;
    case "text":
      button = (
        <Button
          variant="contained"
          onClick={logic}
          style={{
            ...buttonStyle,
            borderRadius: "25px",
            width: "150px",
            height: "50px",
          }}
        >
          <Typography
            style={{ fontSize: "1rem", flex: 1, textTransform: "none" }}
          >
            {text}
          </Typography>
        </Button>
      );
      break;
  }
  return button;
};

const LoginPage = () => {
  const t = useTranslations("LoginPage");
  const [emailText, setEmailText] = useState("");
  const [passwordText, setPasswordText] = useState("");
  const [loginError, setLoginError] = useState<string | null>(null);

  const handleEmailChange = (event: ChangeEvent<HTMLInputElement>) => {
    setEmailText(event.target.value);
    setLoginError(null);
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    setPasswordText(event.target.value);
    setLoginError(null);
  };

  const handleGuestAccount = () => {
    setEmailText("test@email.com");
    setPasswordText("test-password");
    handleLogin();
  }

  const handleLogin = async () => {
    try {
      await axios
        .post("/api/v1/users/login", {
          email: emailText,
          password: passwordText,
        })
        .then((result: any) => {
          if (result.status === 200) {
            const accessToken = result.data.access_token;
            sessionStorage.setItem("access_token", accessToken);
            window.location.href = "/";
          }
        });
    } catch (error: any) {
      setLoginError("Invalid email or password. Please try again.");
    }
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        backgroundColor: "#438EF3",
        height: "100vh",
        color: "white",
      }}
    >
      {/* Language Select */}
      <div
        style={{
          display: "flex",
          marginTop: "15%",
          alignSelf: "flex-end",
          marginRight: "5%",
        }}
      >
        <Typography color="#CCCCCC">
          <Link
            underline="none"
            color={t("lang") === "korean" ? "white" : "inherit"}
            href="/kr/login"
          >
            한국어
          </Link>
          {" | "}
          <Link
            underline="none"
            color={t("lang") === "english" ? "white" : "inherit"}
            href="/en/login"
          >
            English
          </Link>
        </Typography>
      </div>

      {/* Title */}
      <Typography
        style={{
          fontSize: "1.6rem",
          fontWeight: "bold",
          marginTop: "20%",
        }}
      >
        {t("login_title")}
      </Typography>
      <Image
        src={`/sso/${t("login_logo")}.svg`}
        alt="LoginEnglishTitle"
        width={280}
        height={65}
      />

      {/* Body */}
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          marginTop: "20%",
          width: "100%",
          height: "70%",
          borderRadius: "50px 50px 0 0",
          backgroundColor: "white",
        }}
      >
        {/* Manual Login */}
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            marginTop: "8%",
            width: "70%",
          }}
        >
          <TextField
            label={t("email_text")}
            variant="outlined"
            fullWidth
            style={{ marginBottom: "10px" }}
            onChange={handleEmailChange}
          />
          <TextField
            label={t("password_text")}
            variant="outlined"
            fullWidth
            type="password"
            style={{ marginBottom: "5px" }}
            onChange={handlePasswordChange}
          />
          <Typography color="red" fontSize="0.7rem" style={{ height: "20px" }}>
            {loginError}
          </Typography>
          <Button
            style={{
              width: "160px",
              height: "40px",
              borderRadius: "25px",
              backgroundColor: DefaultButtonColor,
              color: "inherit",
            }}
            onClick={handleLogin}
          >
            {t("login_button")}
          </Button>
        </div>

        {/* Divider */}
        <div
          style={{
            display: "flex",
            width: "70%",
            justifyContent: "center",
            margin: "6% 0px",
          }}
        >
          <Divider
            sx={{
              width: "25%",
              borderColor: "#333333",
              margin: "8px 16px",
            }}
          />
          <Typography color="#333333" fontSize="12px">
            {t("or_divider")}
          </Typography>
          <Divider
            sx={{
              width: "25%",
              borderColor: "#333333",
              margin: "8px 16px",
            }}
          />
        </div>

        {/* SSO Button Group */}
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            marginBottom: "10%",
            width: "100%",
          }}
        >
          <SSOButton image="kakao" backgroundColor="#FDDC3F" />
          <SSOButton image="google" backgroundColor="white" />
          <SSOButton text={t("guest_login")} backgroundColor="white" type="text" logic={handleGuestAccount}/>
        </div>

        {/* Footer Links Row */}
        <div
          style={{
            width: "220px",
            display: "flex",
            justifyContent: "space-evenly",
            textAlign: "center",
            fontSize: "0.9rem",
            color: LoginLinkColor,
            textDecoration: "none",
          }}
        >
          <div
            style={{ marginRight: "20px" }}
            onClick={() => {
              window.location.href = "/register";
            }}
          >
            Register
          </div>
          <div>{"|"}</div>
          <div>Forgot Password?</div>
          <div>{"|"}</div>
          <div>Find Email</div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
