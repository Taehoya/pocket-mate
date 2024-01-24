import { Typography, IconButton, Grid } from "@mui/material";

// Icons
import CalendarMonthRoundedIcon from "@mui/icons-material/CalendarMonthRounded";
import InsertChartRoundedIcon from "@mui/icons-material/InsertChartRounded";
import AddCardRoundedIcon from "@mui/icons-material/AddCardRounded";
import DescriptionRoundedIcon from "@mui/icons-material/DescriptionRounded";

// Constants
import { UnactiveColor, DefaultButtonColor } from "../../constants";

interface TripFooterProps {
  onSelect: (componentNumber: number) => void;
  activeComponent: number;
}

const TripFooter: React.FC<TripFooterProps> = ({
  onSelect,
  activeComponent,
}) => {
  const buttonStyle = {
    display: "flex",
    flexDirection: "column",
  };

  const activeButtonStyle = {
    ...buttonStyle,
    color: DefaultButtonColor,
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        borderTop: `1px solid ${UnactiveColor}`,
      }}
    >
      <Grid container justifyContent="space-between" style={{ height: "100%" }}>
        <Grid
          item
          style={{
            flexGrow: 1,
            display: "flex",
            justifyContent: "center",
          }}
        >
          <IconButton
            style={activeComponent === 1 ? activeButtonStyle : buttonStyle}
            onClick={() => onSelect(1)}
          >
            <DescriptionRoundedIcon style={{ fontSize: "35px" }} />
            <Typography fontSize="0.7rem">Trip Info</Typography>
          </IconButton>
        </Grid>
        <Grid
          item
          style={{ flexGrow: 1, display: "flex", justifyContent: "center" }}
        >
          <IconButton
            style={activeComponent === 2 ? activeButtonStyle : buttonStyle}
            onClick={() => onSelect(2)}
          >
            <CalendarMonthRoundedIcon style={{ fontSize: "35px" }} />
            <Typography fontSize="0.7rem">Plan</Typography>
          </IconButton>
        </Grid>
        <Grid
          item
          style={{ flexGrow: 1, display: "flex", justifyContent: "center" }}
        >
          <IconButton
            style={activeComponent === 3 ? activeButtonStyle : buttonStyle}
            onClick={() => onSelect(3)}
          >
            <AddCardRoundedIcon style={{ fontSize: "35px" }} />
            <Typography fontSize="0.7rem">Expense</Typography>
          </IconButton>
        </Grid>
        <Grid
          item
          style={{ flexGrow: 1, display: "flex", justifyContent: "center" }}
        >
          <IconButton
            style={activeComponent === 4 ? activeButtonStyle : buttonStyle}
            onClick={() => onSelect(4)}
          >
            <InsertChartRoundedIcon style={{ fontSize: "35px" }} />
            <Typography fontSize="0.7rem">Graph</Typography>
          </IconButton>
        </Grid>
      </Grid>
    </div>
  );
};

export default TripFooter;
