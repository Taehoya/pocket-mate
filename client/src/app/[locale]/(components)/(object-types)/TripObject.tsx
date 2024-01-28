type TripObject = {
  id: number;
  budget: number;
  countryProperty: {
    id: number,
    code: string,
    name: string,
    currency: string,
  },
  description: string;
  startDateTime: string;
  endDateTime: string;
  name: string;
  noteProperty: {
    bound: string,
    boundColor: string,
    noteColor: string
  }
};

export default TripObject;
