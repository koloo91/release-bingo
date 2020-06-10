import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Entry } from '../models/entry';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  constructor(private httpClient: HttpClient) {
  }

  login(name: string, password: string): Observable<any> {
    const basicAuthHeaderValue = window.btoa(`${name}:${password}`)
    const httpHeaders = new HttpHeaders({
      'Authorization': `Basic ${basicAuthHeaderValue}`
    })
    return this.httpClient.post(`${environment.host}/admin/login`, {}, {headers: httpHeaders});
  }

  getEntries(): Observable<Entry[]> {
    return this.httpClient.get<{ data: Entry[] }>(`${environment.host}/admin/entries`)
      .pipe(
        map(wrapper => wrapper.data)
      );
  }

  createEntry(text: string): Observable<Entry> {
    return this.httpClient.post<Entry>(`${environment.host}/admin/entries`, {text});
  }

  deleteEntry(id: string): Observable<any> {
    return this.httpClient.delete(`${environment.host}/admin/entries/${id}`);
  }

  updateEntry(id: string, text: string, checked: boolean): Observable<any> {
    return this.httpClient.put(`${environment.host}/admin/entries/${id}`, {text, checked});
  }
}
