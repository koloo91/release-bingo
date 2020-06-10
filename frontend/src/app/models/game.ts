export interface Game {
  users: string[];
  user: User;
}

export interface User {
  name: string;
  card: Card;
}

export interface Card {
  rowOne: BingoField[];
  rowTwo: BingoField[];
  rowThree: BingoField[];
  rowFour: BingoField[];
}

export interface BingoField {
  id: string;
  text: string;
  checked: boolean;
}

export interface MarkedEntry {
  id: string;
  checked: boolean;
}
