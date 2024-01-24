import { useState } from "react";
import {
  List,
  ListItem,
  ListItemText,
  Collapse,
  IconButton,
  Divider,
} from "@mui/material";

// ICONS
import ArrowForward from "@mui/icons-material/KeyboardArrowRightRounded";
import ExpandMore from "@mui/icons-material/ExpandMoreRounded";

const AllTransactions = () => {
  const [openDropdowns, setOpenDropdowns] = useState<number[]>([]);

  // Handle dropdown toggle
  const handleDropdownToggle = (index: number) => {
    setOpenDropdowns((prev) => {
      if (prev.includes(index)) {
        // If dropdown is open, close it
        return prev.filter((item) => item !== index);
      } else {
        // If dropdown is closed, open it
        return [...prev, index];
      }
    });
  };

  // Check if a dropdown is open
  const isDropdownOpen = (index: number) => openDropdowns.includes(index);

  return (
    <List>
      <ListItem button onClick={() => handleDropdownToggle(0)}>
        <IconButton style={{ color: "black" }}>
          {isDropdownOpen(0) ? <ExpandMore /> : <ArrowForward />}
        </IconButton>
        <ListItemText primary="Preparation" />
      </ListItem>
      <Collapse in={openDropdowns.includes(0)} timeout="auto" unmountOnExit>
        <List component="div" disablePadding>
          <ListItem button>
            <ListItemText primary="Item 1" />
          </ListItem>
          <ListItem button>
            <ListItemText primary="Item 2" />
          </ListItem>
        </List>
      </Collapse>

      <Divider variant="fullWidth" component="li" />

      <ListItem button onClick={() => handleDropdownToggle(1)}>
        <IconButton style={{ color: "black" }}>
          {isDropdownOpen(0) ? <ExpandMore /> : <ArrowForward />}
        </IconButton>
        <ListItemText primary="Day 1" />
      </ListItem>
      <Collapse in={openDropdowns.includes(1)} timeout="auto" unmountOnExit>
        <List component="div" disablePadding>
          <ListItem button>
            <ListItemText primary="Item 3" />
          </ListItem>
          <ListItem button>
            <ListItemText primary="Item 4" />
          </ListItem>
        </List>
      </Collapse>
    </List>
  );
};

export default AllTransactions;
