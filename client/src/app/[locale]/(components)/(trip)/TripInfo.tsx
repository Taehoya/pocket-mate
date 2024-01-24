"use client";

import { useState, useRef, useEffect } from "react";
import { Box, Button, Typography, IconButton, TextField } from "@mui/material";
import { styled } from "@mui/system";
import Image from "next/image";

// ICONS
import ModeEditIcon from "@mui/icons-material/ModeEdit";

// CONSTANTS
import { DefaultButtonColor, HomeBackgroundColor } from "../../constants";

interface EditFieldProps {
  fieldType?: string;
  fieldTitle: string;
}

const NoBorderTextField = styled(TextField)({
  "& .MuiOutlinedInput-input": {
    fontSize: "1.5rem",
    fontWeight: "bold",
    margin: "1% 0",
    padding: 0,
    width: "auto",
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
});

const EditField: React.FC<EditFieldProps> = ({
  fieldType = "TextField",
  fieldTitle,
}) => {
  const [editMode, setEditMode] = useState(true);
  const textFieldRef = useRef<HTMLInputElement>(null);
  const [noteImage, setNoteImage] = useState("");
  let inputComponent;

  useEffect(() => {
    if (textFieldRef.current) {
      textFieldRef.current.focus();
      // Set the selection range to the end of the input value
      const length = textFieldRef.current.value.length;
      textFieldRef.current.setSelectionRange(length, length);
    }
  }, [editMode]);

  const handleEdit = () => {
    setEditMode(false);
  };

  const handleBlur = () => {
    setEditMode(true);
  };

  const handleNoteImage = (noteImage: string) => {
    setNoteImage(noteImage);
  };

  switch (fieldType) {
    case "TextField":
      inputComponent = (
        <Box sx={{ display: "flex" }}>
          <NoBorderTextField
            onBlur={handleBlur}
            InputProps={{
              readOnly: editMode,
            }}
            inputRef={textFieldRef}
            defaultValue="Hello world"
          />
          <IconButton onClick={handleEdit}>
            <ModeEditIcon />
          </IconButton>
        </Box>
      );
      break;
    case "NoteField":
      inputComponent = (
        <Box sx={{ display: "flex" }}>
          {/* Basic Note */}
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              marginTop: "5%",
              marginBottom: "2%",
            }}
          >
            <Image
              src={`/create-trip/${
                noteImage == "BasicNote" ? "BasicNoteFilled" : "BasicNoteBlank"
              }.svg`}
              alt="My Image"
              width={180}
              height={180}
              onClick={() => handleNoteImage("BasicNote")}
            />
            <Typography sx={{ fontSize: "1rem", marginTop: "5px" }}>
              Basic Note
            </Typography>
          </Box>

          {/* Spring Note */}
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              marginTop: "5%",
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
            <Typography sx={{ fontSize: "1rem", marginTop: "5px" }}>
              Spring Note
            </Typography>
          </Box>
        </Box>
      );
      break;
    case "ColorField":
      inputComponent = (
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
          }}
        >
          {/* Color 1 */}
          <Box
            sx={{
              display: "flex",
              margin: "2% 0",
            }}
          >
            <Typography fontSize="1.5rem" fontWeight="bold" sx={{ flex: 1 }}>
              Binder
            </Typography>
            <Box
              sx={{
                flex: 1,
                backgroundColor: "blue",
                width: "30%",
                borderRadius: "5px",
              }}
            />
            <Box sx={{ flex: 1, marginLeft: "5%" }}>
              <IconButton>
                <ModeEditIcon />
              </IconButton>
            </Box>
          </Box>
          {/* Color 2 */}
          <Box sx={{ display: "flex" }}>
            <Typography fontSize="1.5rem" fontWeight="bold" sx={{ flex: 1 }}>
              Body
            </Typography>
            <Box
              sx={{
                flex: 1,
                backgroundColor: "blue",
                width: "30%",
                borderRadius: "5px",
              }}
            />
            <Box sx={{ flex: 1, marginLeft: "5%" }}>
              <IconButton>
                <ModeEditIcon />
              </IconButton>
            </Box>
          </Box>
        </Box>
      );
      break;
  }

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        marginTop: "5%",
        marginLeft: "3%",
      }}
    >
      <Typography fontSize="0.9rem">{fieldTitle}</Typography>
      {inputComponent}
    </Box>
  );
};

const TripInfo = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      <EditField fieldTitle="Trip Name" />
      <EditField fieldTitle="Travel To" />
      <EditField fieldTitle="Start Date" />
      <EditField fieldTitle="End Date" />
      <EditField fieldTitle="Currency Unit" />
      <EditField fieldTitle="Travel Note Type" fieldType="NoteField" />
      <EditField fieldTitle="Travel Note Color" fieldType="ColorField" />

      {/* Button */}
      <Button
        style={{
          marginTop: "5%",
          marginBottom: "5%",
          borderRadius: 27,
          height: "5.5%",
          width: "70%",
          fontSize: "1rem",
          fontWeight: "bold",
          backgroundColor: DefaultButtonColor,
          color: "white",
        }}
      >
        Complete
      </Button>
    </div>
  );
};

export default TripInfo;
