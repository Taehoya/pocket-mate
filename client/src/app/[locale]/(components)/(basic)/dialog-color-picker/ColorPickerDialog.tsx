import { useState } from "react";
import {
  Avatar,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
  Typography,
} from "@mui/material";

// Icons
import CheckIcon from "@mui/icons-material/Check";

interface ColorPickerDialogProps {
  open: boolean;
  onClose: () => void;
}

const ColorPickerDialog: React.FC<ColorPickerDialogProps> = ({
  open,
  onClose,
}) => {
  const [currentColor, setCurrentColor] = useState("#7195BD");

  const handleColorChange = (value: string) => {
    setCurrentColor(value);
  };

  const handleConfirm = () => {
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle style={{ textAlign: "center" }}>
        Select Binder Color
      </DialogTitle>
      <DialogContent>
        <Grid container justifyContent="center" gap={3.5}>
          {shadesArray.map((value, index) => (
            <Grid key={index} item xs={5}>
              <div
              onClick={() => handleColorChange(value)}
                style={{
                  height: "50px",
                  backgroundColor: value,
                  borderRadius: "10px",
                  position: "relative",
                }}
              >
                {currentColor === value && (
                  <Avatar
                    style={{
                      position: "absolute",
                      top: "50%",
                      left: "50%",
                      transform: "translate(-50%, -50%)",
                      backgroundColor: "white",
                      width: "35px",
                      height: "35px",
                    }}
                  >
                    <CheckIcon
                      style={{ color: value, width: "30px", height: "30px" }}
                    />
                  </Avatar>
                )}
              </div>
            </Grid>
          ))}
        </Grid>
      </DialogContent>
      <DialogActions
        style={{
          justifyContent: "center",
          borderTop: "1px solid #ccc",
          padding: "16px",
        }}
      >
        <Button onClick={handleConfirm} color="primary">
          <Typography color="black" fontWeight={400} fontSize="18px">
            OK
          </Typography>
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ColorPickerDialog;

const shadesArray: string[] = [
  "#E5EEF5",
  "#C8D8E5",
  "#ABC2D6",
  "#8EABCC",
  "#7195BD",
  "#548EAF",
  "#3778A0",
  "#1A6292",
  "#004C83",
  "#00354B",
];
