import React, { useEffect, useRef, useState } from "react";
import { IconButton, TextareaAutosize } from "@mui/material";
import { styled } from "@mui/styles";

// ICONS
import ModeEditIcon from "@mui/icons-material/ModeEdit";

const NoBorderTextField = styled(TextareaAutosize)({
  fontSize: "20px",
  fontWeight: 500,
  margin: "1% 0",
  padding: 0,
  width: "250",
  border: "none",
  outline: "none",
  letterSpacing: "-0.5px",
});

interface TripTextFieldProps {
  text: string;
  active?: boolean;
}

const TripTextField: React.FC<TripTextFieldProps> = ({
  text,
  active = false,
}) => {
  const [editMode, setEditMode] = useState(true);
  const textFieldRef = useRef<HTMLTextAreaElement>(null);

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

  return (
    <div style={{ display: "flex" }}>
      <NoBorderTextField
        onBlur={handleBlur}
        readOnly={editMode}
        ref={textFieldRef}
        defaultValue={text}
      />
      {active && (
        <IconButton onClick={handleEdit}>
          <ModeEditIcon />
        </IconButton>
      )}
    </div>
  );
};

export default TripTextField;
