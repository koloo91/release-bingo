import { ActionReducerMap, MetaReducer } from '@ngrx/store';
import { environment } from '../../environments/environment';
import { GameState, gameReducer } from './game/game.reducer';
import { adminReducer, AdminState } from './admin/admin.reducer';

export interface AppState {
  gameState: GameState;
  adminState: AdminState;
}

export const reducers: ActionReducerMap<AppState> = {
  gameState: gameReducer,
  adminState: adminReducer
};


export const metaReducers: MetaReducer<AppState>[] = !environment.production ? [] : [];
