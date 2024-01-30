import { useState } from "react";
import { Typography } from "@mui/material";
import Image from "next/image";

const TripNoteField = () => {
  const [noteImage, setNoteImage] = useState("BasicNote");

  const handleNoteImage = (noteImage: string) => {
    setNoteImage(noteImage);
  };

  return (
    <div style={{ display: "flex" }}>
      {/* Basic Note */}
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          marginTop: "5%",
          marginBottom: "2%",
        }}
      >
        <Image
          src={`/create-trip/${
            noteImage == "BasicNote" ? "BasicNoteFilled" : "BasicNoteBlank"
          }.svg`}
          alt="My Image"
          width={180}
          height={180}
          onClick={() => handleNoteImage("BasicNote")}
        />
        <Typography sx={{ fontSize: "1rem", marginTop: "5px" }}>
          Basic Note
        </Typography>
      </div>

      {/* Spring Note */}
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          marginTop: "5%",
        }}
      >
        <Image
          src={`/create-trip/${
            noteImage == "SpringNote" ? "SpringNoteFilled" : "SpringNoteBlank"
          }.svg`}
          alt="My Image"
          width={180}
          height={180}
          onClick={() => handleNoteImage("SpringNote")}
        />
        <Typography sx={{ fontSize: "1rem", marginTop: "5px" }}>
          Spring Note
        </Typography>
      </div>
    </div>
  );
};

export default TripNoteField;
