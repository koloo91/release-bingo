import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, map, switchMap } from 'rxjs/operators';
import { of } from 'rxjs';
import { AdminService } from '../../services/admin.service';
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
  loadEntriesSuccessAction,
  updateEntryAction,
  updateEntryFailedAction,
  updateEntrySuccessAction
} from './admin.actions';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable()
export class AdminEffects {

  adminLogin$ = createEffect(() => this.actions$.pipe(
    ofType(adminLoginAction),
    switchMap(({name, password}) => {
      return this.adminService.login(name, password)
        .pipe(
          map(() => adminLoginSuccessAction({name, password})),
          catchError(error => {
            console.log(error);
            return of(adminLoginFailedAction({error: {message: 'Invalid credentials'}}))
          })
        );
    })
  ));

  loadEntries$ = createEffect(() => this.actions$.pipe(
    ofType(loadEntriesAction),
    switchMap(() => this.adminService.getEntries()
      .pipe(
        map(entries => loadEntriesSuccessAction({entries})),
        catchError(error => of(loadEntriesFailedAction({error: error.error})))
      ))
  ));

  createEntry$ = createEffect(() => this.actions$.pipe(
    ofType(createEntryAction),
    switchMap(({text}) => this.adminService.createEntry(text)
      .pipe(
        map(() => createEntrySuccessAction()),
        catchError((error: HttpErrorResponse) => of(createEntryFailedAction({error: {message: error.message}})))
      ))
  ))

  updateEntry$ = createEffect(() => this.actions$.pipe(
    ofType(updateEntryAction),
    switchMap(({id, text, checked}) => this.adminService.updateEntry(id, text, checked)
      .pipe(
        map(() => updateEntrySuccessAction()),
        catchError((error: HttpErrorResponse) => of(updateEntryFailedAction({error: {message: error.message}})))
      ))
  ))

  deleteEntry$ = createEffect(() => this.actions$.pipe(
    ofType(deleteEntryAction),
    switchMap(({id}) => this.adminService.deleteEntry(id)
      .pipe(
        map(() => deleteEntrySuccessAction()),
        catchError((error: HttpErrorResponse) => of(deleteEntryFailedAction({error: {message: error.message}})))
      ))
  ))

  constructor(private actions$: Actions,
              private adminService: AdminService) {
  }
}
