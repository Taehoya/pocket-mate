import React from "react";
import {
  Avatar,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Typography,
} from "@mui/material";

// Icons
import HotelIcon from "@mui/icons-material/LocalHotelOutlined";

interface TransactionItemProps {
    amount: string;
    name: string;
}

const TransactionItem: React.FC<TransactionItemProps> = ({amount, name}) => {
  return (
    <ListItem button>
      <ListItemAvatar>
        <Avatar style={{ borderRadius: "10px", width: "48px", height: "48px" }}>
          <HotelIcon />
        </Avatar>
      </ListItemAvatar>
      <ListItemText
        primary={
          <Typography
            style={{ fontWeight: 400, fontSize: "22px" }}
          >{amount}
          </Typography>
        }
        secondary={name}
        style={{ marginLeft: "5px" }}
      />
    </ListItem>
  );
};

export default TransactionItem;
