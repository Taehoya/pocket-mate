import { Button, Typography } from "@mui/material";

// Constants
import { DefaultButtonColor, UnactiveColor } from "@/app/[locale]/constants";

const TransactionTemplate = () => {
  return (
    <div style={{ display: "flex", flexDirection: "column", height: "100%" }}>
      {/* Amount Field  */}
      <div
        style={{
          marginTop: "5%",
          display: "flex",
          flexDirection: "column",
          background: "#F9F9F9",
          height: "15%",
        }}
      ></div>

      {/* Bubble Dropdown Section */}
      <div
        style={{
          display: "flex",
          height: "66px",
          padding: "0px 2%",
          margin: "10% 0px",
        }}
      >
        {/* Date Selection */}
        <div
          style={{
            flex: 2,
            height: "100%",
            borderRadius: "10px",
            border: `1px solid ${UnactiveColor}`,
            margin: "0px 2%",
          }}
        ></div>

        {/* Payment Type */}
        <div
          style={{
            flex: 1,
            height: "100%",
            borderRadius: "10px",
            border: `1px solid ${UnactiveColor}`,
            margin: "0px 2%",
          }}
        ></div>
      </div>

      {/* Expense Name */}
      <div style={{ height: "45px", margin: "5% 0px" }}></div>

      {/* Category List*/}

      {/* Add Button */}
      <div style={{ display: "flex", justifyContent: "center" }}>
        <Button
          style={{
            height: "50px",
            width: "350px",
            flexShrink: 0,
            borderRadius: "25px",
            backgroundColor: DefaultButtonColor,
            boxShadow: "0px 4px 4px 0px rgba(0, 0, 0, 0.25)"
          }}
        >
          <Typography
            color="white"
            fontSize="20px"
            fontWeight={800}
            textAlign="center"
          >
            Add Expense
          </Typography>
        </Button>
      </div>
    </div>
  );
};

export default TransactionTemplate;
