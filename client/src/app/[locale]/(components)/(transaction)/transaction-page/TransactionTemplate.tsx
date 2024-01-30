import { TextField } from "@mui/material";

import DefaultButton from "../../(basic)/default-button/DefaultButton";
import CategoryList from "../transaction-component/category-list/CategoryList";
import DropdownSection from "../transaction-component/dropdown-bubble/DropdownSection";
import InputAmount from "../transaction-component/input-amount/InputAmount";

const TransactionTemplate = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        width: "100%",
        height: "100%",
      }}
    >
      {/* Amount Field  */}
      <InputAmount amount={500} />

      {/* Bubble Dropdown Section */}
      <DropdownSection />

      {/* Expense Name Field */}
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
          marginBottom: "100px",
        }}
      >
        <DefaultButton name="Add Expense" />
      </div>
    </div>
  );
};

export default TransactionTemplate;
