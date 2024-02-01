import { Button, Typography } from "@mui/material";

// Constants
import { DefaultButtonColor, UnactiveButtonColor } from "@/app/[locale]/constants";

interface DefaultButtonProps {
  name: string;
  active?: boolean;
  onClick?: () => void;
}

const DefaultButton: React.FC<DefaultButtonProps> = ({
  name,
  onClick,
  active = true,
}) => {
  return (
    <Button
      style={{
        height: "50px",
        width: "350px",
        flexShrink: 0,
        borderRadius: "25px",
        backgroundColor: active ? DefaultButtonColor : UnactiveButtonColor,
        boxShadow: "0px 4px 4px 0px rgba(0, 0, 0, 0.25)",
      }}
      onClick={onClick}
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
