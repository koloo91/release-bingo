import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AppState } from '../reducers';
import { Store } from '@ngrx/store';
import { selectLastAdminError } from '../reducers/admin/admin.selectors';
import { filter } from 'rxjs/operators';
import { selectLastGameError } from '../reducers/game/game.selectors';

@Injectable({
  providedIn: 'root'
})
export class GameErrorService {

  constructor(store: Store<AppState>, snackbar: MatSnackBar) {
    store.select(selectLastGameError)
      .pipe(
        filter(err => !!err)
      )
      .subscribe(httpError => {
        snackbar.open(httpError.message, 'OK', {
          duration: 5000
        });
      })
  }
}
