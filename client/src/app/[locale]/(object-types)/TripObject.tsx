type TripObject = {
  id: number;
  budget: number;
  countryID: number;
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
