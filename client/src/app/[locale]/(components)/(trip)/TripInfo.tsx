"use client";

import React, { useState } from "react";
import { Box, Typography } from "@mui/material";
import DefaultButton from "../(basic)/default-button/DefaultButton";
import TripColorField from "./tripinfo-component/TripColorField";
import TripNoteField from "./tripinfo-component/TripNoteField";
import TripTextField from "./tripinfo-component/TripTextField";

// Icons
import EditNoteIcon from "@mui/icons-material/EditNote";

// Constants
import { DefaultButtonColor } from "../../constants";

interface EditFieldProps {
  fieldTitle: string;
  children: React.ReactNode;
}

const EditField: React.FC<EditFieldProps> = ({ fieldTitle, children }) => {
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
      {children}
    </Box>
  );
};

const TripInfo = () => {
  const [toggleEdit, setToggleEdit] = useState(false);

  const handleToggleEdit = () => {
    setToggleEdit((prev) => !prev);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        position: "relative",
      }}
    >
      {/* Toggle Edit Button */}
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          width: "20%",
          height: "3%",
          position: "absolute",
          top: "0.5%",
          right: "2%",
          borderRadius: "15px",
          border: toggleEdit ? "none" : "1px solid #999",
          color: toggleEdit ? "white" : "black",
          backgroundColor: toggleEdit ? DefaultButtonColor : "white",
        }}
        onClick={handleToggleEdit}
      >
        {toggleEdit ? (
          <EditNoteIcon />
        ) : (
          <Typography fontWeight={600} fontSize="18px">
            Edit
          </Typography>
        )}
      </div>

      <EditField fieldTitle="Trip Name">
        <TripTextField
          active={toggleEdit}
          text="My own healing trip test test test 22221"
        />
      </EditField>
      <EditField fieldTitle="Travel To">
        <TripTextField active={toggleEdit} text="My own healing trip" />
      </EditField>
      <EditField fieldTitle="Currency Unit">
        <TripTextField active={toggleEdit} text="My own healing trip" />
      </EditField>
      <EditField fieldTitle="Travel Note Type">
        <TripNoteField />
      </EditField>
      <EditField fieldTitle="Travel Note Color">
        <TripColorField active={toggleEdit} />
      </EditField>

      {/* Button */}
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          width: "100%",
          margin: "10% 0%",
        }}
      >
        <DefaultButton name="Complete" active={toggleEdit} />
      </div>
    </div>
  );
};

export default TripInfo;
