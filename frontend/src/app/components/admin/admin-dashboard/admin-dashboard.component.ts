import { Component, OnInit } from '@angular/core';
import { AppState } from '../../../reducers';
import { Store } from '@ngrx/store';
import { deleteEntryAction, loadEntriesAction, updateEntryAction } from '../../../reducers/admin/admin.actions';
import {
  selectEntries,
  selectEntryDeletedSuccessful,
  selectIsLoadingEntries
} from '../../../reducers/admin/admin.selectors';
import { Observable } from 'rxjs';
import { Entry } from '../../../models/entry';
import { MatDialog } from '@angular/material/dialog';
import { AddEntryDialogComponent } from '../add-entry-dialog/add-entry-dialog.component';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrls: ['./admin-dashboard.component.scss']
})
export class AdminDashboardComponent implements OnInit {

  isLoadingEntries$: Observable<boolean>;
  entries$: Observable<Entry[]>;

  constructor(private store: Store<AppState>,
              public dialog: MatDialog) {
    this.isLoadingEntries$ = store.select(selectIsLoadingEntries);
    this.entries$ = store.select(selectEntries);
    this.store.select(selectEntryDeletedSuccessful)
      .pipe(filter(f => f))
      .subscribe(() => this.store.dispatch(loadEntriesAction()))
  }

  ngOnInit(): void {
    this.store.dispatch(loadEntriesAction());
  }

  openAddEntryDialog() {
    const dialogRef = this.dialog.open(AddEntryDialogComponent, {
      width: '500px'
    });

    dialogRef.afterClosed().subscribe(result => {
      if (!result.canceled) {
        this.store.dispatch(loadEntriesAction());
      }
    })
  }

  onCheckBoxChanged(event: MatCheckboxChange, entry: Entry) {
    this.store.dispatch(updateEntryAction({id: entry.id, text: entry.text, checked: event.checked}));
  }

  deleteEntry(id: string) {
    this.store.dispatch(deleteEntryAction({id}));
  }
}
