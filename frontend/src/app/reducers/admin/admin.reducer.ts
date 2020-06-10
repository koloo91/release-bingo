import { Action, createReducer, on } from '@ngrx/store';
import {
  adminLoginAction,
  adminLoginFailedAction,
  adminLoginSuccessAction,
  createEntryAction,
  createEntryFailedAction,
  createEntrySuccessAction,
  deleteEntryAction, deleteEntryFailedAction,
  deleteEntrySuccessAction,
  loadEntriesAction,
  loadEntriesFailedAction,
  loadEntriesSuccessAction
} from './admin.actions';
import { HttpError } from '../../models/http_error';
import { Entry } from '../../models/entry';

export interface AdminState {
  isLoggingIn: boolean;
  authorizationHeader?: string;
  lastHttpError?: HttpError;
  isLoadingEntries: boolean;
  isCreatingEntry: boolean;
  isDeletingEntry: boolean;
  entryCreatedSuccessful: boolean;
  entryDeletedSuccessful: boolean;
  entries: Entry[];
}

export const initialState: AdminState = {
  isLoggingIn: false,
  authorizationHeader: null,
  lastHttpError: null,
  isLoadingEntries: false,
  isCreatingEntry: false,
  isDeletingEntry: false,
  entryCreatedSuccessful: false,
  entryDeletedSuccessful: false,
  entries: []
}

const reducer = createReducer(
  initialState,
  on(adminLoginAction, state => {
    return {
      ...state,
      isLoggingIn: true
    }
  }),
  on(adminLoginSuccessAction, (state, {name, password}) => {
    return {
      ...state,
      authorizationHeader: 'Basic ' + window.btoa(`${name}:${password}`),
      isLoggingIn: false,
      lastHttpError: null
    }
  }),
  on(adminLoginFailedAction, (state, {error}) => {
    return {
      ...state,
      authorizationHeader: null,
      isLoggingIn: false,
      lastHttpError: error
    }
  }),
  on(loadEntriesAction, state => {
    return {
      ...state,
      isLoadingEntries: true
    }
  }),
  on(loadEntriesSuccessAction, (state, {entries}) => {
    return {
      ...state,
      isLoadingEntries: false,
      entries: entries,
      entryCreatedSuccessful: false,
      entryDeletedSuccessful: false
    }
  }),
  on(loadEntriesFailedAction, (state, {error}) => {
    return {
      ...state,
      isLoadingEntries: false,
      lastHttpError: error
    }
  }),
  on(createEntryAction, state => {
    return {
      ...state,
      isCreatingEntry: true,
      entryCreatedSuccessful: false,
      lastHttpError: null
    }
  }),
  on(createEntrySuccessAction, state => {
    return {
      ...state,
      isCreatingEntry: false,
      entryCreatedSuccessful: true,
      lastHttpError: null
    }
  }),
  on(createEntryFailedAction, (state, {error}) => {
    return {
      ...state,
      isCreatingEntry: false,
      lastHttpError: error
    }
  }),
  on(deleteEntryAction, state => {
    return {
      ...state,
      isDeletingEntry: true,
      entryDeletedSuccessful: false,
      lastHttpError: null
    }
  }),
  on(deleteEntrySuccessAction, state => {
    return {
      ...state,
      isDeletingEntry: false,
      entryDeletedSuccessful: true,
      lastHttpError: null
    }
  }),
  on(deleteEntryFailedAction, (state, {error}) => {
    return {
      ...state,
      isDeletingEntry: false,
      entryDeletedSuccessful: false,
      lastHttpError: error
    }
  })
)

export function adminReducer(state: AdminState | undefined, action: Action) {
  return reducer(state, action)
}
