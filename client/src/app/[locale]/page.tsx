"use client";

import React, { useState, useEffect } from "react";
import MultiPageForm from "./(components)/(extra)/MultiPageForm";
import SwipableCards from "./(components)/(home)/SwipableCards";
import GridCards from "./(components)/(home)/GridCards";
import {
  Button,
  Dialog,
  IconButton,
  MenuItem,
  Select,
  SelectChangeEvent,
  Typography,
} from "@mui/material";
import TripObject from "./(components)/(object-types)/TripObject";
import axios from "axios";
import { useTranslations } from "next-intl";
import DefaultButton from "./(components)/(basic)/default-button/DefaultButton";
import WebWrapper from "./(wrapper)/WebWrapper";

// ICONS
import SettingsIcon from "@mui/icons-material/Settings";
import WindowIcon from "@mui/icons-material/Window";
import ListIcon from "@mui/icons-material/FormatListBulleted";

// CONSTANTS
import {
  BackgroundColor,
  DefaultButtonColor,
  HomeBackgroundColor,
} from "./constants";

export default function Home() {
  const t = useTranslations("HomePage");
  const [isFormOpen, setIsFormOpen] = useState(false);
  const [dropdownValue, setDropdownValue] = useState("current");
  const [swipeView, setSwipeView] = useState(true);
  const [pastTripList, setPastTripList] = useState<TripObject[] | undefined>();
  const [futureTripList, setFutureTripList] = useState<
    TripObject[] | undefined
  >();
  const [currentTripList, setCurrentTripList] = useState<
    TripObject[] | undefined
  >();
  const [tripList, setTripList] = useState<TripObject[]>();
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    const accessToken = sessionStorage.getItem("access_token");
    if (!accessToken) window.location.href = "/login";
    else {
      fetchTrips(accessToken);
    }
  }, []);

  const fetchTrips = async (token: String) => {
    setLoading(true);
    const trips = await axios.get("/api/v1/trips", {
      headers: { Authorization: `Bearer ${token}` },
    });
    setPastTripList(trips.data.past);
    setFutureTripList(trips.data.future);
    setCurrentTripList(trips.data.current);
    setTripList(trips.data.current);
    setLoading(false);
  };

  const addTravelNote = () => {
    setIsFormOpen(true);
  };

  const closeTravelNote = () => {
    setIsFormOpen(false);
  };

  const controlChangeView = () => {
    setSwipeView(!swipeView);
  };

  const handleDropdown = (event: SelectChangeEvent<string>) => {
    const time = event.target.value as string;
    setDropdownValue(time);
    if (time === "past") setTripList(pastTripList);
    else if (time === "coming-soon") setTripList(futureTripList);
    else setTripList(currentTripList);
  };

  return (
    <WebWrapper>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          height: "100vh",
          backgroundColor: HomeBackgroundColor,
        }}
      >
        {/* Header */}
        <div
          style={{
            marginTop: "3%",
            textAlign: "center",
            flex: 2,
          }}
        >
          {/* Title Text */}
          <Typography style={{ fontSize: "1.1rem" }}>{t("header")}</Typography>

          {/* Button Row */}
          <div
            style={{
              marginTop: "7%",
              padding: "0px 2%",
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            {/* Settings */}
            <IconButton>
              <SettingsIcon />
            </IconButton>

            {/* Dropdown */}
            <Select
              value={dropdownValue}
              onChange={handleDropdown}
              style={{
                width: "190px",
                height: "40px",
                borderRadius: "25px",
                backgroundColor: "white",
                color: "black",
                boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.2)",
                border: "none",
              }}
            >
              <MenuItem value="current">{t("current")}</MenuItem>
              <MenuItem value="coming-soon">{t("future")}</MenuItem>
              <MenuItem value="past">{t("present")}</MenuItem>
            </Select>

            {/* Change View */}
            <IconButton onClick={controlChangeView}>
              {swipeView ? <WindowIcon /> : <ListIcon />}
            </IconButton>
          </div>
        </div>

        {/* Body */}
        {!loading && (
          <div style={{ flex: 8 }}>
            {swipeView ? (
              <SwipableCards trips={tripList} />
            ) : (
              <GridCards trips={tripList} />
            )}
          </div>
        )}

        {/* Add Travel Button */}
        <div
          style={{
            flex: 1,
            display: "flex",
            flexDirection: "column",
            justifyContent: "flex-end",
            alignItems: "center",
            width: "100%",
            paddingBottom: "15%",
          }}
        >
          <DefaultButton name={t("add_note")} onClick={addTravelNote}/>
        </div>

        {/* New Note Modal */}
        <Dialog
          open={isFormOpen}
          fullScreen
          onClose={closeTravelNote}
          PaperProps={{
            style: {
              overflowY: "hidden",
              maxHeight: "100%",
              backgroundColor: BackgroundColor,
            },
          }}
        >
          <MultiPageForm closeForm={closeTravelNote} />
        </Dialog>
      </div>
    </WebWrapper>
  );
}
