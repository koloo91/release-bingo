import { Action, createReducer, on, State } from '@ngrx/store';
import { connectToGameAction, onGameConnectionError, onGameStateChangedAction } from './game.actions';
import { User } from '../../models/game';
import { HttpError } from '../../models/http_error';

export interface GameState {
  isConnectingToGame: boolean;
  users: string[];
  user?: User;
  lastError?: HttpError
}

export const initialState: GameState = {
  isConnectingToGame: false,
  users: [],
  user: null,
  lastError: null
}

const reducer = createReducer(
  initialState,
  on(connectToGameAction, state => {
    return {
      ...state,
      isConnectingToGame: true,
      lastError: null
    }
  }),
  on(onGameStateChangedAction, (state, {game}) => {
    return {
      ...state,
      isConnectingToGame: false,
      users: game.users,
      user: game.user,
      lastError: null
    }
  }),
  on(onGameConnectionError, (state, {error}) => {
    return {
      ...state,
      isConnectingToGame: false,
      lastError: error
    }
  })
)

export function gameReducer(state: GameState | undefined, action: Action) {
  return reducer(state, action)
}
