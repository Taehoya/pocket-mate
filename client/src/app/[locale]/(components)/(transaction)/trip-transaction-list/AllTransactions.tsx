import React from "react";
import { List, Divider } from "@mui/material";

// ICONS
import DayTransactions from "./DayTransactions";

const AllTransactions = () => {
  return (
    <List>
      {daysList.map((value, index) => (
        <React.Fragment key={index}>
          <DayTransactions title={value} />
          {index < daysList.length - 1 && (
            <Divider variant="fullWidth" component="li" key={`divider-${index}`} />
          )}
        </React.Fragment>
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
