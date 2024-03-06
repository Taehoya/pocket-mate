type TripByID = {
  id: number;
  budget: number;
  countryProperty: {
    id: number;
    code: string;
    name: string;
    currency: string;
  };
  description: string;
  startDateTime: string;
  endDateTime: string;
  name: string;
  noteProperty: {
    boundColor: string;
    noteColor: string;
    noteType: string;
  };
  top5Transactions?: Array<{
    amount: number;
    category: {
      id: number;
      name: string;
    };
    description: string;
    name: string;
    transactionDateTime: string;
  }>;
  totalExpense?: number;
  transactions?: Array<{
    amount: number;
    category: {
      id: number;
      name: string;
    };
    description: string;
    name: string;
    transactionDateTime: string;
  }>;
};

export default TripByID;
