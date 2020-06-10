import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AppState } from '../../../reducers';
import { Store } from '@ngrx/store';
import { adminLoginAction } from '../../../reducers/admin/admin.actions';
import { selectAdminIsLoggedIn } from '../../../reducers/admin/admin.selectors';
import { filter } from 'rxjs/operators';
import { Router } from '@angular/router';

@Component({
  selector: 'app-admin-login',
  templateUrl: './admin-login.component.html',
  styleUrls: ['./admin-login.component.scss']
})
export class AdminLoginComponent implements OnInit {

  loginFormGroup = new FormGroup({
    name: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required])
  })

  constructor(private store: Store<AppState>, router: Router) {
    store.select(selectAdminIsLoggedIn)
      .pipe(
        filter(_ => _)
      )
      .subscribe(() => {
        router.navigate(['admin/dashboard']);
      })
  }

  ngOnInit(): void {
  }

  login() {
    const {name, password} = this.loginFormGroup.value;
    this.store.dispatch(adminLoginAction({name, password}));
  }
}
