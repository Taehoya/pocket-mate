import React, { useState } from "react";
import {
  Collapse,
  IconButton,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";
import TransactionItem from "./TransactionItem";

// ICONS
import ArrowForward from "@mui/icons-material/KeyboardArrowRightRounded";
import ExpandMore from "@mui/icons-material/ExpandMoreRounded";

interface DayTransactionsProps {
  title: string;
}

const DayTransactions: React.FC<DayTransactionsProps> = ({ title }) => {
  const [openDropdown, setOpenDropdown] = useState(false);

  // Handle dropdown toggle
  const handleDropdownToggle = () => {
    setOpenDropdown((prev) => !prev);
  };

  return (
    <>
      <ListItem button onClick={handleDropdownToggle}>
        <IconButton style={{ color: "black" }}>
          {openDropdown ? <ExpandMore /> : <ArrowForward />}
        </IconButton>
        <ListItemText primary={title} />
      </ListItem>
      <Collapse in={openDropdown} timeout="auto" unmountOnExit>
        <List component="div" disablePadding style={{ width: "100%" }}>
          <TransactionItem amount="$450" name="Hilton Hotel"/>
          <TransactionItem  amount="$1250" name="Royal Dragon Inn"/>
          <TransactionItem  amount="$680" name="The Macabre"/>
        </List>
      </Collapse>
    </>
  );
};

export default DayTransactions;
