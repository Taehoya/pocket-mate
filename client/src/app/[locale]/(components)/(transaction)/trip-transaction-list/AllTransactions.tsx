import { List, Divider } from "@mui/material";

// ICONS
import DayTransactions from "./DayTransactions";

const AllTransactions = () => {
  return (
    <List>
      {daysList.map((value) => (
        <>
          <DayTransactions title={value} />
          <Divider variant="fullWidth" component="li" />
        </>
      ))}
    </List>
  );
};

export default AllTransactions;

const daysList: string[] = [
  "Preparation",
  "Day 1",
  "Day 2",
  "Day 3",
  "Day 4",
  "Day 5",
];
