import DropdownBubble from "./DropdownBubble";

const DropdownSection = () => {
  return (
    <div
      style={{
        display: "flex",
        minHeight: "8.5%",
        padding: "0px 2%",
        margin: "10% 0px",
      }}
    >
      {/* Date Dropdown */}
      <div
        style={{
          flex: 3,
          height: "100%",
          margin: "0px 2%",
        }}
      >
        <DropdownBubble label="Select Dates" content={daysList}/>
      </div>

      {/* Payment Type Dropdown */}
      <div
        style={{
          flex: 2,
          height: "100%",
          margin: "0px 2%",
        }}
      >
        <DropdownBubble label="Payment" content={paymentList}/>
      </div>
    </div>
  );
};

export default DropdownSection;

const daysList: string[] = ["Travel Preparation", "Day 1", "Day 2", "Day 3", "Day 4", "Day 5"];
const paymentList: string[] = ["Card", "Cash"];
