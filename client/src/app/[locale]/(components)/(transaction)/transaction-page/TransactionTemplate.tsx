import { Button, TextField, Typography } from "@mui/material";

// Constants
import { DefaultButtonColor, UnactiveColor } from "@/app/[locale]/constants";
import CategoryList from "../../(basic)/category-list/CategoryList";
import DefaultButton from "../../(basic)/default-button/DefaultButton";

const TransactionTemplate = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        width: "100%",
        height: "100vh",
      }}
    >
      {/* Amount Field  */}
      <div
        style={{
          marginTop: "5%",
          display: "flex",
          flexDirection: "column",
          background: "#F9F9F9",
          height: "150px",
        }}
      >
        {/* Curreny Toggle */}

        {/* Amount Input Value */}

        {/* Converted Currency Value */}
      </div>

      {/* Bubble Dropdown Section */}
      <div
        style={{
          display: "flex",
          minHeight: "10%",
          padding: "0px 2%",
          margin: "10% 0px",
        }}
      >
        {/* Date Dropdown */}
        <div
          style={{
            flex: 2,
            height: "100%",
            borderRadius: "10px",
            border: `1px solid ${UnactiveColor}`,
            margin: "0px 2%",
          }}
        ></div>

        {/* Payment Type Dropdown */}
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
      <div
        style={{
          height: "95px",
          width: "100%",
          margin: "5% 0px",
          display: "flex",
          justifyContent: "center",
        }}
      >
        <div style={{ width: "90%" }}>
          <TextField
            fullWidth
            id="standard-input-name"
            label="Expense Name"
            variant="standard"
            InputLabelProps={{
              shrink: true,
              style: { fontSize: "20px" },
            }}
            InputProps={{
              style: { height: "45px" },
            }}
            placeholder="Please enter name of expense"
          />
        </div>
      </div>

      {/* Category List*/}
      <CategoryList />

      {/* Add Button */}
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          width: "100%",
          marginBottom: "10%",
        }}
      >
        <DefaultButton name="Add Expense" />
      </div>
    </div>
  );
};

export default TransactionTemplate;
