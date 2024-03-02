import { useEffect, useState } from "react";
import { format } from "date-fns";
import { DateRange, DayPicker } from "react-day-picker";
import { enUS, ko } from "date-fns/locale";
import { useTranslations } from "next-intl";
import "react-day-picker/dist/style.css";

interface DatePickerProps {
  handleDateChange: Function;
  range?: DateRange | undefined;
}

const DatePicker: React.FC<DatePickerProps> = ({ handleDateChange, range }) => {
  const t = useTranslations("DatePicker");
  const pastMonth = range.from;
  const [language, setLanguage] = useState(enUS);

  useEffect(() => {
    switch (t("language")) {
      case "korean":
        setLanguage(ko);
        break;
      case "english":
        setLanguage(enUS);
        break;
      default:
        setLanguage(enUS);
        break;
    }
  }, []);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        // backgroundColor: "green",
      }}
    >
      <DayPicker
        id="date-picker"
        mode="range"
        locale={language}
        defaultMonth={pastMonth}
        selected={range}
        onSelect={(event: DateRange) => {
          let result = event;
          if (!event) result = { from: range.from, to: range.from };
          else if (!event.to) result = { from: event.from, to: event.from };
          handleDateChange(result);
        }}
        modifiersStyles={{
          selected: {
            backgroundColor: "#FDDC3F",
            color: "#333333",
          },
        }}
        // styles={{
        //   root: { backgroundColor: "pink", width: "100%", height: "100%", justifyContent: "center" },
        //   caption: { backgroundColor: "blue", width: "100%" },
        //   table: { backgroundColor: "blue", width: "100%" },
        //   head: {
        //     backgroundColor: "red",
        //     width: "80px",
        //     height: "50px",
        //   },
        //   day: {
        //     width: "50px",
        //     height: "50px",
        //   },
        // }}
      />
    </div>
  );
};

export default DatePicker;
