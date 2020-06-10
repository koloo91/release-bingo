import { createAction, props } from '@ngrx/store';
import { Game } from '../../models/game';
import { HttpError } from '../../models/http_error';


export const connectToGameAction = createAction('[Game] Connect to game', props<{ name: string }>());

export const onGameStateChangedAction = createAction('[Game] Game state changed', props<{ game: Game }>());

export const onGameConnectionError = createAction('[Game] Game connection error', props<{error: HttpError}>());
