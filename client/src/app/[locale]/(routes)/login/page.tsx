"use client";

import React from "react";
import Image from "next/image";
import { useTranslations } from "next-intl";
import { Divider, Typography, Button, Link, TextField } from "@mui/material";

// Constants
import { DefaultButtonColor, LoginLinkColor } from "@/app/[locale]/constants";

interface SSOButtonProps {
  image: string;
  backgroundColor: string;
  text: string;
}

const SSOButton: React.FC<SSOButtonProps> = ({
  image,
  backgroundColor,
  text,
}) => {
  const imageSrc = `/sso/${image}.svg`;

  return (
    <Button
      variant="contained"
      style={{
        backgroundColor: backgroundColor,
        color: "black",
        borderRadius: "5px",
        width: "280px",
        height: "50px",
        display: "flex",
        alignItems: "center",
        padding: "8px",
        justifyContent: "flex-start",
        paddingLeft: "20px",
        marginBottom: "7%",
      }}
      startIcon={<Image src={imageSrc} alt="SSO Icon" width="45" height="45" />}
    >
      <Typography style={{ fontSize: "1rem", flex: 1, textTransform: "none" }}>
        {text}
      </Typography>
    </Button>
  );
};

const LoginPage = () => {
  const t = useTranslations("LoginPage");

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
            marginTop: "12%",
            width: "70%",
          }}
        >
          <TextField
            label={t("email_text")}
            variant="outlined"
            fullWidth
            style={{ marginBottom: "10px" }}
          />
          <TextField
            label={t("password_text")}
            variant="outlined"
            fullWidth
            style={{ marginBottom: "20px" }}
          />
          <Button
            style={{
              width: "160px",
              height: "40px",
              borderRadius: "25px",
              backgroundColor: DefaultButtonColor,
              color: "inherit",
            }}
          >
            {t("login_button")}
          </Button>
        </div>

        {/* Divider */}
        <div
          style={{ display: "flex", width: "70%", justifyContent: "center", margin: "6% 0px" }}
        >
          <Divider
            sx={{
              width: "25%",
              borderColor: "#333333",
              margin: "8px 16px",
            }}
          />
          <Typography color="#333333" fontSize="12px">
            간편 로그인
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
            flexDirection: "column",
            justifyContent: "center",
          }}
        >
          <SSOButton
            image="kakao"
            text={t("kakao_login")}
            backgroundColor="#FDDC3F"
          />
          <SSOButton
            image="google"
            text={t("google_login")}
            backgroundColor="white"
          />
        </div>

        {/* Registration Row */}
        <div
          style={{
            width: "220px",
            display: "flex",
            justifyContent: "center",
            textAlign: "center",
            fontSize: "1rem",
            color: LoginLinkColor,
            textDecoration: "none",
          }}
        >
          <a href="/register" style={{ marginRight: "20px" }}>
            Register
          </a>
          <div>{"|"}</div>
          <a href="/login" style={{ marginLeft: "20px" }}>
            Login
          </a>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
