import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { AppState } from '../../../reducers';
import { connectToGameAction } from '../../../reducers/game/game.actions';
import { selectIsConnectedToGame, selectIsConnectingToGame } from '../../../reducers/game/game.selectors';
import { Observable } from 'rxjs';
import { filter } from 'rxjs/operators';
import { Router } from '@angular/router';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.scss']
})
export class WelcomeComponent implements OnInit {

  isConnectingToGame$: Observable<boolean>;

  nameFormGroup = new FormGroup({
    name: new FormControl('', [Validators.required])
  })

  constructor(private store: Store<AppState>, router: Router) {
    this.isConnectingToGame$ = store.select(selectIsConnectingToGame)
    this.isConnectingToGame$.subscribe(value => {
      if (value) {
        this.nameFormGroup.controls['name'].disable();
      } else {
        this.nameFormGroup.controls['name'].enable();
      }
    });
    store.select(selectIsConnectedToGame)
      .pipe(filter(f => f))
      .subscribe(() => router.navigate(['current-game']));
  }

  ngOnInit(): void {
  }

  onSubmit() {
    console.log('onSubmit');
    if (!this.nameFormGroup.valid) {
      return;
    }
    console.log('onSubmit isValid');
    this.store.dispatch(connectToGameAction({name: this.nameFormGroup.value.name}))
  }
}
