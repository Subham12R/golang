export interface Movie {
  id: string;
  title: string;
  genre: string;
  duration: number;
}

export interface Seat {
  row: number;
  column: number;
  seatId: string;
  status: 'available' | 'booked' | 'reserved';
}