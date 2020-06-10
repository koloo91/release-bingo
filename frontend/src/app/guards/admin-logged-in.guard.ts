import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router } from '@angular/router';
import { Observable, of } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { selectAdminIsLoggedIn } from '../reducers/admin/admin.selectors';
import { first, map, mergeMap, take } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AdminLoggedInGuard implements CanActivate {
  constructor(private store: Store<AppState>,
              private router: Router) {
  }

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    return this.store.select(selectAdminIsLoggedIn)
      .pipe(
        take(1),
        mergeMap((loggedIn) => {
          if (!loggedIn) {
            console.log('not logged in. redirecting');
            this.router.navigate(['admin/login']);
            return of(false);
          }
          return of(true);
        })
      );
  }

}
