import { Component, OnInit } from '@angular/core';
import { State, Store } from '@ngrx/store';
import { AppState } from '../../../reducers';
import { selectCard, selectUserName, selectUsers } from '../../../reducers/game/game.selectors';
import { Observable } from 'rxjs';
import { Card } from '../../../models/game';

@Component({
  selector: 'app-current-game',
  templateUrl: './current-game.component.html',
  styleUrls: ['./current-game.component.scss']
})
export class CurrentGameComponent implements OnInit {

  users$: Observable<string[]>;
  userName$: Observable<string>;
  card$: Observable<Card>;

  constructor(store: Store<AppState>) {
    this.users$ = store.select(selectUsers);
    this.userName$ = store.select(selectUserName);
    this.card$ = store.select(selectCard);
  }

  ngOnInit(): void {
  }

}
