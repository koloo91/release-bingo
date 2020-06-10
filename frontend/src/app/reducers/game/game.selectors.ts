import { AppState } from '../index';
import { createSelector } from '@ngrx/store';

export const gameFeature = (state: AppState) => state.gameState;

export const selectUsers = createSelector(
  gameFeature,
  state => state.users
);

export const selectIsConnectingToGame = createSelector(
  gameFeature,
  state => state.isConnectingToGame
);

export const selectIsConnectedToGame = createSelector(
  gameFeature,
  state => !!state.user
);

export const selectUserName = createSelector(
  gameFeature,
  state => state.user.name
);

export const selectCard = createSelector(
  gameFeature,
  state => state.user.card
);

export const selectLastGameError = createSelector(
  gameFeature,
  state => state.lastError
)
