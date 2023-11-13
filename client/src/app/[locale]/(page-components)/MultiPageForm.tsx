"use client";

import React, { useState, useEffect, ReactNode, ChangeEvent } from "react";
import {
  Stepper,
  Step,
  Button,
  IconButton,
  Paper,
  LinearProgress,
  Typography,
  TextField,
  Grid,
} from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import CloseIcon from "@mui/icons-material/Close";
import { DateRange } from "react-date-range";
import { addDays } from "date-fns";
import "react-date-range/dist/styles.css";
import "react-date-range/dist/theme/default.css";
import { useTranslations } from "next-intl";

import Image from "next/image";

// CONSTANTS
import { DefaultButtonColor } from "../constants";

const steps = ["Step 0", "Step 1", "Step 2", "Step 3", "Step 4", "step 5"];

interface StepPageProps {
  children: ReactNode;
  buttonClick: () => void;
  isDisable: Boolean;
}

interface MultiPageFormProps {
  closeForm: () => void;
}

const maxCharacters: number = 24;

const styles = {
  resultListStyle: {
    borderBottom: "2px dashed grey",
    padding: "15px 0px",
  },
  resultKeyStyle: {
    color: "grey",
    fontSize: "1 rem",
    display: "flex",
    justifyContent: "flex-start",
  },
  resultValueStyle: {
    fontSize: "0.95 rem",
    display: "flex",
    justifyContent: "flex-end",
  },
};

const StepPage: React.FC<StepPageProps> = ({
  children,
  buttonClick,
  isDisable = true,
}) => {
  const t = useTranslations("TripCreation");

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        height: "100%",
        width: "100%",
        marginTop: "30px",
      }}
    >
      <div style={{ width: "90%" }}>{children}</div>
      <div
        style={{
          flex: 1,
          display: "flex",
          flexDirection: "column",
          justifyContent: "flex-end",
          alignItems: "center",
          width: "100%",
          marginBottom: "20%",
        }}
      >
        <Button
          onClick={buttonClick}
          disabled={isDisable}
          style={{
            backgroundColor: isDisable ? "grey" : DefaultButtonColor,
            color: "white",
            borderRadius: "25px",
            width: "90%",
            padding: "10px 0px",
          }}
        >
          {t("button_next")}
        </Button>
      </div>
    </div>
  );
};

const MultiPageForm: React.FC<MultiPageFormProps> = ({ closeForm }) => {
  const [activeStep, setActiveStep] = useState(0);
  const [progress, setProgress] = useState(0);
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(addDays(new Date(), 7));
  const [title, setTitle] = useState("");
  const [destination, setDestination] = useState("");
  const [noteImage, setNoteImage] = useState("");
  const t = useTranslations("TripCreation");

  useEffect(() => {
    setProgress((activeStep / (steps.length - 1)) * 100);
  }, [activeStep]);

  const handleTitle = (event: ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value;
    if (newValue.length <= maxCharacters) {
      setTitle(newValue);
    }
  };

  const handleDateChange = (ranges: {
    selection: { startDate: Date; endDate: Date };
  }) => {
    const { selection } = ranges;
    setStartDate(selection.startDate);
    setEndDate(selection.endDate);
  };

  const handleDestination = (event: ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value;
    setDestination(newValue);
  };

  const handleNoteImage = (noteImage: string) => {
    setNoteImage(noteImage);
  };

  const handleNext = () => {
    setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100vh",
      }}
    >
      <Stepper
        activeStep={activeStep}
        alternativeLabel
        style={{
          backgroundColor: "transparent",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <IconButton disabled={activeStep === 0} onClick={handleBack}>
          <ArrowBackIcon />
        </IconButton>
        <Typography style={{ fontSize: "1rem" }}>Add a travel note</Typography>
        <IconButton onClick={closeForm}>
          <CloseIcon />
        </IconButton>
      </Stepper>
      <LinearProgress
        variant="determinate"
        value={progress}
        classes={{ bar: { bacgroundColor: "#0e85ff" } }}
        style={{
          backgroundColor: "#E6E0E9",
          height: "5px",
          borderRadius: "10px",
          width: "90%",
          margin: "0px 20px",
          marginTop: "20px",
          marginBottom: "0px",
        }}
      />
      <div
        style={{
          flex: 1,
          display: "flex",
          justifyContent: "center",
        }}
      >
        {activeStep === 0 && (
          <StepPage buttonClick={handleNext} isDisable={false}>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}
            >
              <div
                style={{
                  fontSize: "1.3rem",
                  width: "95%",
                  marginBottom: "5px",
                }}
              >
                {t("welcome")}
              </div>
              <div
                style={{ fontSize: "1.7rem", fontWeight: "bold", width: "95%" }}
              >
                {t("welcome_sub")}
              </div>
              <div style={{ marginTop: "30%", marginBottom: "30px" }}>
                <Image
                  src="/create-trip/CreateTripStart.png"
                  alt="StartTripImage"
                  width={270}
                  height={220}
                />
              </div>
            </div>
          </StepPage>
        )}
        {activeStep === 1 && (
          <StepPage
            buttonClick={handleNext}
            isDisable={!(destination.length > 0)}
          >
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
                width: "80%",
                lineHeight: 1.4,
              }}
            >
              {t("destination_title")}
            </div>
            <TextField
              variant="standard"
              fullWidth
              placeholder={t("destination_input")}
              style={{
                marginTop: "25px",
              }}
              value={destination}
              onChange={handleDestination}
            />
          </StepPage>
        )}
        {activeStep === 2 && (
          <StepPage buttonClick={handleNext} isDisable={false}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
                width: "80%",
              }}
            >
              {t("date_title")}
            </div>
            <div style={{ marginTop: "30px" }}>
              <DateRange
                ranges={[
                  {
                    startDate: startDate,
                    endDate: endDate,
                    key: "selection",
                  },
                ]}
                onChange={handleDateChange}
                editableDateInputs={true}
                moveRangeOnFirstSelection={false}
                showDateDisplay={false}
                // rangeColors={["yellow", "#3ecf8e", "#fed14c"]}
              />
            </div>
          </StepPage>
        )}
        {activeStep === 3 && (
          <StepPage buttonClick={handleNext} isDisable={!(title.length > 0)}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
              }}
            >
              {t("trip_title")}
            </div>
            <div>
              <TextField
                variant="standard"
                fullWidth
                placeholder={t("trip_input")}
                style={{
                  marginTop: "20px",
                }}
                value={title}
                onChange={handleTitle}
              />
              <div
                style={{
                  color: "grey",
                  marginTop: "5px",
                  fontSize: "0.8rem",
                }}
              >
                {title.length} / {maxCharacters} {t("character")}
              </div>
            </div>
          </StepPage>
        )}
        {activeStep === 4 && (
          <StepPage buttonClick={handleNext} isDisable={!noteImage}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
              }}
            >
              {t("note_type_title")}
            </div>
            <div
              style={{
                display: "flex",
                marginTop: "35%",
              }}
            >
              <div
                style={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Image
                  src={`/create-trip/${
                    noteImage == "BasicNote"
                      ? "BasicNoteFilled"
                      : "BasicNoteBlank"
                  }.svg`}
                  alt="My Image"
                  width={180}
                  height={180}
                  onClick={() => handleNoteImage("BasicNote")}
                />
                <Typography style={{ fontSize: "1rem", marginTop: "5px" }}>
                  Basic Note
                </Typography>
              </div>
              <div
                style={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Image
                  src={`/create-trip/${
                    noteImage == "SpringNote"
                      ? "SpringNoteFilled"
                      : "SpringNoteBlank"
                  }.svg`}
                  alt="My Image"
                  width={180}
                  height={180}
                  onClick={() => handleNoteImage("SpringNote")}
                />
                <Typography style={{ fontSize: "1rem", marginTop: "5px" }}>
                  Spring Note
                </Typography>
              </div>
            </div>
          </StepPage>
        )}
        {activeStep === 5 && (
          <StepPage buttonClick={() => {}} isDisable={false}>
            <div
              style={{
                display: "flex",
                flexDirection: "column",
                textAlign: "center",
                alignItems: "center",
              }}
            >
              <div
                style={{ fontSize: "2rem", fontWeight: "bold", width: "80%" }}
              >
                {t("final_title")}
              </div>
              <div style={{ marginTop: "20px", marginBottom: "30px" }}>
                <Image
                  src="/NewTripPlaceholder.png"
                  alt="My Image"
                  width={200}
                  height={200}
                />
              </div>

              <Grid container>
                {/* Trip Title */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    {t("container_title")}
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {title}
                  </Grid>
                </Grid>

                {/* Trip Destination */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    {t("container_destination")}
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {destination}
                  </Grid>
                </Grid>

                {/* Trip Start Date */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    {t("container_start_date")}
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {startDate?.toDateString()}
                  </Grid>
                </Grid>

                {/* Trip End Date */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    {t("container_end_date")}
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {endDate?.toDateString()}
                  </Grid>
                </Grid>
              </Grid>
            </div>
          </StepPage>
        )}
      </div>
    </div>
  );
};

export default MultiPageForm;
