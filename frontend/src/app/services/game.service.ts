import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { environment } from '../../environments/environment';
import { Game } from '../models/game';
import { Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { EMPTY, Observable } from 'rxjs';
import { onGameConnectionError, onGameStateChangedAction } from '../reducers/game/game.actions';

@Injectable({
  providedIn: 'root'
})
export class GameService {
  private webSocketSubject: WebSocketSubject<Game>;

  constructor(private state: Store<AppState>) {
  }

  getWsHost(): string {
    return environment.wsHost || window.location.host;
  }

  getWsUrl(): string {
    const protocol = window.location.protocol === 'http:' ? 'ws:' : 'wss:';
    return `${protocol}//${this.getWsHost()}`;
  }

  connectToGame(name: string): Observable<any> {
    name = encodeURIComponent(name)
    this.webSocketSubject = webSocket<Game>(`${this.getWsUrl()}/bingo/${name}`);
    this.webSocketSubject.subscribe((game) => {
        console.log(game);
        this.state.dispatch(onGameStateChangedAction({game}))
      },
      (error) => {
        console.log(error);
        this.state.dispatch(onGameConnectionError({error: {message: 'Benutzername ist bereits vergeben'}}))
      });
    return EMPTY;
  }
}
