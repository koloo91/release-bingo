import { AppState } from '../index';
import { createSelector } from '@ngrx/store';

export const adminFeature = (state: AppState) => state.adminState;

export const selectLastAdminError = createSelector(
  adminFeature,
  state => state.lastHttpError
);

export const selectAdminIsLoggedIn = createSelector(
  adminFeature,
  state => !!state.authorizationHeader
);

export const selectIsLoadingEntries = createSelector(
  adminFeature,
  state => state.isLoadingEntries
);

export const selectEntries = createSelector(
  adminFeature,
  state => state.entries
);

export const selectIsCreatingEntry = createSelector(
  adminFeature,
  state => state.isCreatingEntry
);

export const selectEntryCreatedSuccessful = createSelector(
  adminFeature,
  state => state.entryCreatedSuccessful
);

export const selectEntryDeletedSuccessful = createSelector(
  adminFeature,
  state => state.entryDeletedSuccessful
);
