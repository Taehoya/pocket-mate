"use client";

import React from "react";
import { Typography, Button } from "@mui/material";
import Image from "next/image";
import { LoginLinkColor } from "@/app/[locale]/constants";

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
  const imageSrc = `/sso/${image}.png`;

  return (
    <Button
      variant="contained"
      style={{
        backgroundColor: backgroundColor,
        color: "black",
        borderRadius: "5px",
        width: "230px",
        display: "flex",
        alignItems: "center",
        padding: "8px",
        justifyContent: "flex-start",
        paddingLeft: "20px",
        marginBottom: "7%",
      }}
      startIcon={<Image src={imageSrc} alt="SSO Icon" width="24" height="24" />}
    >
      <Typography style={{ fontSize: "0.8rem", flex: 1 }} variant="button">
        {text}
      </Typography>
    </Button>
  );
};

const SSOLoginPage = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Typography
        style={{
          fontSize: "1.6rem",
          fontWeight: "bold",
          marginTop: "40%",
          color: "black",
        }}
      >
        Let's make a trip together,
      </Typography>
      <Image src="/logo.png" alt="StartTripImage" width={190} height={190} />

      {/* SSO Button Group */}
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          marginTop: "30%",
        }}
      >
        <SSOButton
          image="naver"
          text="continue with naver"
          backgroundColor="#FDDC3F"
        />
        <SSOButton
          image="google"
          text="continue with google"
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
          marginTop: "5%",
          fontSize: "0.8rem",
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
  );
};

export default SSOLoginPage;
