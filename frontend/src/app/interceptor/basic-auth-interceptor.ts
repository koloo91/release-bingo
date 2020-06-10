import { Injectable } from '@angular/core';
import { HttpEvent, HttpHandler, HttpHeaders, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Store } from '@ngrx/store';
import { AppState } from '../reducers';
import { switchMap, take } from 'rxjs/operators';

@Injectable()
export class BasicAuthInterceptor implements HttpInterceptor {
  constructor(private store: Store<AppState>) {
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return this.store.select(state => state.adminState.authorizationHeader)
      .pipe(
        take(1),
        switchMap(headerValue => {
          console.log(headerValue);
          if (headerValue) {
            const request = req.clone({
              headers: new HttpHeaders({
                'Authorization': headerValue
              })
            });
            return next.handle(request);
          }
          return next.handle(req);
        }),
      );
  }

}
