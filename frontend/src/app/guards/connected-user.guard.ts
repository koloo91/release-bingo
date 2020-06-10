import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router } from '@angular/router';
import { Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { selectIsConnectedToGame } from '../reducers/game/game.selectors';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ConnectedUserGuard implements CanActivate {
  constructor(private store: Store<AppState>,
              private router: Router) {
  }

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    return this.store.select(selectIsConnectedToGame)
      .pipe(
        map((isConnected) => {
          if (!isConnected) {
            this.router.navigate(['']);
            return false;
          }
          return true;
        }));
  }

}
