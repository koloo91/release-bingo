import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { WelcomeComponent } from './components/game/welcome/welcome.component';
import { CurrentGameComponent } from './components/game/current-game/current-game.component';
import { ConnectedUserGuard } from './guards/connected-user.guard';
import { AdminLoginComponent } from './components/admin/admin-login/admin-login.component';
import { AdminDashboardComponent } from './components/admin/admin-dashboard/admin-dashboard.component';
import { AdminLoggedInGuard } from './guards/admin-logged-in.guard';


const routes: Routes = [
  {
    path: '',
    component: WelcomeComponent
  },
  {
    path: 'current-game',
    component: CurrentGameComponent,
    canActivate: [ConnectedUserGuard]
  },
  {
    path: 'admin',
    children: [
      {
        path: 'login',
        component: AdminLoginComponent
      },
      {
        path: 'dashboard',
        component: AdminDashboardComponent,
        canActivate: [AdminLoggedInGuard],
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
