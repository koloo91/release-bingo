import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AdminService } from '../../../services/admin.service';
import { Store } from '@ngrx/store';
import { AppState } from '../../../reducers';
import { selectEntryCreatedSuccessful, selectIsCreatingEntry } from '../../../reducers/admin/admin.selectors';
import { Observable } from 'rxjs';
import { createEntryAction } from '../../../reducers/admin/admin.actions';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-add-entry-dialog',
  templateUrl: './add-entry-dialog.component.html',
  styleUrls: ['./add-entry-dialog.component.scss']
})
export class AddEntryDialogComponent implements OnInit {

  entryFormGroup = new FormGroup({
    text: new FormControl('', [Validators.required])
  });

  isCreatingEntry: Observable<boolean>;

  constructor(public dialogRef: MatDialogRef<AddEntryDialogComponent>,
              private store: Store<AppState>) {
    this.isCreatingEntry = store.select(selectIsCreatingEntry)
    store.select(selectEntryCreatedSuccessful)
      .pipe(filter(f => f))
      .subscribe(() => this.dialogRef.close({canceled: false}))
  }

  ngOnInit(): void {
  }

  onCancelClick() {
    this.dialogRef.close({canceled: true});
  }

  onOkClick() {
    const {text} = this.entryFormGroup.value;
    this.store.dispatch(createEntryAction({text}));
  }
}
