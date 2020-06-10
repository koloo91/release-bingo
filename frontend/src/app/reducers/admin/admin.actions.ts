import { createAction, props } from '@ngrx/store';
import { HttpError } from '../../models/http_error';
import { Entry } from '../../models/entry';


export const adminLoginAction = createAction('[Admin] Login', props<{ name: string, password: string }>());
export const adminLoginSuccessAction = createAction('[Admin] Login success', props<{ name: string, password: string }>());
export const adminLoginFailedAction = createAction('[Admin] Login failed', props<{ error: HttpError }>());

export const loadEntriesAction = createAction('[Admin] Load entries');
export const loadEntriesSuccessAction = createAction('[Admin] Load entries success', props<{ entries: Entry[] }>());
export const loadEntriesFailedAction = createAction('[Admin] Load entries failed', props<{ error: HttpError }>());

export const createEntryAction = createAction('[Admin] Create entry', props<{ text: string }>());
export const createEntrySuccessAction = createAction('[Admin] Create entry success');
export const createEntryFailedAction = createAction('[Admin] Create entry failed', props<{ error: HttpError }>());

export const updateEntryAction = createAction('[Admin] Update entry', props<{ id: string, text: string, checked: boolean }>());
export const updateEntrySuccessAction = createAction('[Admin] Update entry success');
export const updateEntryFailedAction = createAction('[Admin] Update entry failed', props<{ error: HttpError }>());

export const deleteEntryAction = createAction('[Admin] Delete entry', props<{ id: string }>());
export const deleteEntrySuccessAction = createAction('[Admin] Delete entry success');
export const deleteEntryFailedAction = createAction('[Admin] Delete entry failed', props<{ error: HttpError }>());
