import { Player } from "@lottiefiles/react-lottie-player";
import { Typography } from "@mui/material";

const TripGraph = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Player
        autoplay
        loop
        src="https://lottie.host/4fb6060e-e633-41aa-96fd-5ed252a84f74/Y1zq71W96W.json"
        style={{ height: "350px", width: "100%" }}
      />
      <Typography fontSize={20} fontWeight={800}>COMING SOON</Typography>
    </div>
  );
};

export default TripGraph;
