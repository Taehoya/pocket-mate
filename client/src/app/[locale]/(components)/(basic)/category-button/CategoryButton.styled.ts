import { styled } from "@mui/material/styles";
import { Grid, Typography } from "@mui/material";

export const GridMain = styled(Grid)({
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  borderRadius: "10px",
  border: "1.5px solid #FB3F04",
  height: "100px",
  padding: "10px",
});

export const DivIcon = styled("div")({
  borderRadius: "10px",
  backgroundColor: "#FB3F04",
  width: "48px",
  height: "48px",
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
  color: "white",
});

export const IconWord = styled(Typography)({ whiteSpace: "pre-line" });

const style = {
  GridMain: GridMain,
  DivIcon: DivIcon,
  IconWord: IconWord,
};

export default style;
