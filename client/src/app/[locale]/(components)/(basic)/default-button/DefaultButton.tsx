import { Button, Typography } from "@mui/material";

// Constants
import { DefaultButtonColor } from "@/app/[locale]/constants";

interface DefaultButtonProps {
  name: string;
}

const DefaultButton: React.FC<DefaultButtonProps> = ({ name }) => {
  return (
    <Button
      style={{
        height: "50px",
        width: "350px",
        flexShrink: 0,
        borderRadius: "25px",
        backgroundColor: DefaultButtonColor,
        boxShadow: "0px 4px 4px 0px rgba(0, 0, 0, 0.25)",
      }}
    >
      <Typography
        color="white"
        fontSize="17px"
        fontWeight={800}
        textAlign="center"
      >
        {name}
      </Typography>
    </Button>
  );
};

export default DefaultButton;
