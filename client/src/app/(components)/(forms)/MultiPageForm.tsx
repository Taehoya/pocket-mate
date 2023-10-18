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

import Image from "next/image";

const steps = ["Step 1", "Step 2", "Step 3", "Step 4"];

interface StepPageProps {
  children: ReactNode;
  buttonClick: () => void;
  isDisable: Boolean;
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

const StepPage: React.FC<StepPageProps> = ({ children, buttonClick, isDisable = true}) => {
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
            backgroundColor: isDisable ? "grey" : "#fb3f04",
            color: "white",
            borderRadius: "25px",
            width: "90%",
            padding: "10px 0px",
          }}
        >
          Sounds great
        </Button>
      </div>
    </div>
  );
};

const MultiPageForm: React.FC = () => {
  const [activeStep, setActiveStep] = useState(0);
  const [progress, setProgress] = useState(25);
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(addDays(new Date(), 7));
  const [title, setTitle] = useState("");
  const [destination, setDestination] = useState("");

  useEffect(() => {
    setProgress(((activeStep + 1) / steps.length) * 100);
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
        <IconButton onClick={() => {}}>
          <CloseIcon />
        </IconButton>
      </Stepper>
      <LinearProgress
        variant="determinate"
        value={progress}
        color="success"
        style={{
          backgroundColor: "white",
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
          <StepPage buttonClick={handleNext}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
                width: "80%",
                lineHeight: 1.4,
              }}
            >
              Where are you going for the trip?
            </div>
            <TextField
              variant="standard"
              fullWidth
              placeholder="Please choose a travel destination"
              style={{
                marginTop: "25px",
              }}
              value={destination}
              onChange={handleDestination}
            />
          </StepPage>
        )}
        {activeStep === 1 && (
          <StepPage buttonClick={handleNext}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
                width: "80%",
              }}
            >
              When shall we go?
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
        {activeStep === 2 && (
          <StepPage buttonClick={handleNext}>
            <div
              style={{
                fontSize: "1.7rem",
                fontWeight: "bold",
              }}
            >
              What is the Trip's name?
            </div>
            <div>
              <TextField
                variant="standard"
                fullWidth
                placeholder="Propose a travel title"
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
                {title.length} / {maxCharacters} Characters
              </div>
            </div>
          </StepPage>
        )}
        {activeStep === 3 && (
          <StepPage buttonClick={() => {}}>
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
                A new travel note has been created
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
                    Note name
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {title}
                  </Grid>
                </Grid>

                {/* Trip Destination */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    Travel Destination
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {destination}
                  </Grid>
                </Grid>

                {/* Trip Start Date */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    Start Date
                  </Grid>
                  <Grid item xs={6} style={styles.resultValueStyle}>
                    {startDate?.toDateString()}
                  </Grid>
                </Grid>

                {/* Trip End Date */}
                <Grid container style={styles.resultListStyle}>
                  <Grid item xs={6} style={styles.resultKeyStyle}>
                    End Date
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
