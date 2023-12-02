"use client";

import { useState, useRef, useEffect } from "react";
import { Box, Typography, IconButton, TextField } from "@mui/material";
import { styled } from "@mui/system";

// ICONS
import CloseIcon from "@mui/icons-material/Close";
import ModeEditIcon from "@mui/icons-material/ModeEdit";

// CONSTANTS
import { HomeBackgroundColor } from "../constants";

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
    width: "max-content",
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
    // setEditMode(true);
  };

  switch (fieldType) {
    case "TextField":
      inputComponent = (
        <Box sx={{ display: "flex" }} onBlur={handleBlur}>
          <NoBorderTextField
            disabled={editMode}
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

const EditTripPage = () => {
  return (
    <div style={{ position: "relative", backgroundColor: HomeBackgroundColor }}>
      {/* Header */}
      <Box
        sx={{
          width: "100%",
          height: "5%",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          borderBottom: "2px solid #000",
          position: "fixed",
          zIndex: 0,
          backgroundColor: HomeBackgroundColor,
        }}
      >
        <Typography sx={{ flexGrow: 1, textAlign: "center" }}>
          Edit Travel Note
        </Typography>
        <IconButton onClick={() => {}}>
          <CloseIcon />
        </IconButton>
      </Box>

      {/* Body */}
      <Box
        sx={{
          width: "100%",
          display: "flex",
          flexDirection: "column",
          paddingTop: "15%",
          zIndex: -1,
        }}
      >
        <Typography sx={{ flexGrow: 1, textAlign: "center" }}>TETS </Typography>
        <EditField fieldTitle="Trip Name" />
        <EditField fieldTitle="Travel To" />
      </Box>
    </div>
  );
};

export default EditTripPage;
