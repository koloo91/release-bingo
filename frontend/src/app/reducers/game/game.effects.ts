import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { GameService } from '../../services/game.service';
import { connectToGameAction } from './game.actions';
import { catchError, mergeMap } from 'rxjs/operators';
import { EMPTY } from 'rxjs';

@Injectable()
export class GameEffects {

  connectToGame$ = createEffect(() => this.actions$.pipe(
    ofType(connectToGameAction),
    mergeMap(({name}) => {
      console.log(`User ${name} connects to game`);
      return this.gameService.connectToGame(name)
        .pipe(
          catchError(() => EMPTY)
        )
    })
  ));

  constructor(private actions$: Actions,
              private gameService: GameService) {
  }
}
