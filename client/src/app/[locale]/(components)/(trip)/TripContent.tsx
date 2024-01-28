import { useState } from "react";
import { Tab, Tabs, List, Paper } from "@mui/material";
import AllTransactions from "../(transaction)/trip-transaction-list/AllTransactions";

const TripContent = () => {
  const [activeTab, setActiveTab] = useState(0);

  // Handle tab change
  const handleTabChange = (event: any, newValue: number) => {
    setActiveTab(newValue);
  };

  return (
    <div style={{ width: "100%", height: "100%" }}>
      <Tabs
        value={activeTab}
        onChange={handleTabChange}
        aria-label="Tabs with dropdowns"
        sx={{
          "& .MuiTabs-indicator": {
            background: "white",
          },
        }}
      >
        <Tab
          label="Expense List"
          style={{
            backgroundColor: activeTab === 0 ? "#fff" : "#2196f3",
            color: activeTab === 0 ? "#2196f3" : "#fff",
            flexGrow: 1,
          }}
        />
        <Tab
          label="Checklist"
          style={{
            backgroundColor: activeTab === 1 ? "#fff" : "#2196f3",
            color: activeTab === 1 ? "#2196f3" : "#fff",
            flexGrow: 1,
          }}
        />
      </Tabs>

      {activeTab === 0 && <AllTransactions />}

      {activeTab === 1 && <List></List>}
    </div>
  );
};

export default TripContent;
