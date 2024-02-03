import { useState } from "react";
import { Avatar, AvatarGroup, IconButton } from "@mui/material";
import EmailDialog from "../(basic)/dialog-email/EmailDialog";

// Constants
import { DefaultButtonColor } from "../../constants";

// Icons
import GroupAddIcon from "@mui/icons-material/GroupAdd";

const MemberAvatars = () => {
  const [emailDialogOpen, setEmailDialogOpen] = useState(false);

  const handleEmailDialogOpen = () => {
    setEmailDialogOpen(true);
  };

  const handleEmailDialogClose = () => {
    setEmailDialogOpen(false);
  };

  return (
    <div
      style={{
        width: "80%",
        height: "30px",
        marginTop: "10%",
        display: "flex",
        justifyContent: "center",
      }}
    >
      <AvatarGroup total={8}>
        <Avatar alt="Remy Sharp" src="/" sx={{ bgcolor: "green" }} />
        <Avatar alt="Agnes Walker" src="/" sx={{ bgcolor: "lightblue" }} />
        <Avatar alt="Trevor Henderson" src="/" sx={{ bgcolor: "red" }} />
      </AvatarGroup>
      <Avatar sx={{ marginLeft: "3%", backgroundColor: DefaultButtonColor }}>
        <IconButton onClick={handleEmailDialogOpen}>
          <GroupAddIcon />
        </IconButton>
      </Avatar>

      {/* Email Dialog */}
      <EmailDialog open={emailDialogOpen} onClose={handleEmailDialogClose} />
    </div>
  );
};

export default MemberAvatars;
