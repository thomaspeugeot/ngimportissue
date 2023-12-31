// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { FormFieldStringDB } from './formfieldstring-db';
import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormFieldStringService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormFieldStringServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formfieldstringsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.formfieldstringsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formfieldstrings';
  }

  /** GET formfieldstrings from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB[]> {
    return this.getFormFieldStrings(GONG__StackPath, frontRepo)
  }
  getFormFieldStrings(GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormFieldStringDB[]>(this.formfieldstringsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<FormFieldStringDB[]>('getFormFieldStrings', []))
      );
  }

  /** GET formfieldstring by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {
    return this.getFormFieldString(id, GONG__StackPath, frontRepo)
  }
  getFormFieldString(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formfieldstringsUrl}/${id}`;
    return this.http.get<FormFieldStringDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formfieldstring id=${id}`)),
      catchError(this.handleError<FormFieldStringDB>(`getFormFieldString id=${id}`))
    );
  }

  /** POST: add a new formfieldstring to the server */
  post(formfieldstringdb: FormFieldStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {
    return this.postFormFieldString(formfieldstringdb, GONG__StackPath, frontRepo)
  }
  postFormFieldString(formfieldstringdb: FormFieldStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormFieldStringDB>(this.formfieldstringsUrl, formfieldstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formfieldstringdb id=${formfieldstringdb.ID}`)
      }),
      catchError(this.handleError<FormFieldStringDB>('postFormFieldString'))
    );
  }

  /** DELETE: delete the formfieldstringdb from the server */
  delete(formfieldstringdb: FormFieldStringDB | number, GONG__StackPath: string): Observable<FormFieldStringDB> {
    return this.deleteFormFieldString(formfieldstringdb, GONG__StackPath)
  }
  deleteFormFieldString(formfieldstringdb: FormFieldStringDB | number, GONG__StackPath: string): Observable<FormFieldStringDB> {
    const id = typeof formfieldstringdb === 'number' ? formfieldstringdb : formfieldstringdb.ID;
    const url = `${this.formfieldstringsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormFieldStringDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formfieldstringdb id=${id}`)),
      catchError(this.handleError<FormFieldStringDB>('deleteFormFieldString'))
    );
  }

  /** PUT: update the formfieldstringdb on the server */
  update(formfieldstringdb: FormFieldStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {
    return this.updateFormFieldString(formfieldstringdb, GONG__StackPath, frontRepo)
  }
  updateFormFieldString(formfieldstringdb: FormFieldStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<FormFieldStringDB> {
    const id = typeof formfieldstringdb === 'number' ? formfieldstringdb : formfieldstringdb.ID;
    const url = `${this.formfieldstringsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormFieldStringDB>(url, formfieldstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formfieldstringdb id=${formfieldstringdb.ID}`)
      }),
      catchError(this.handleError<FormFieldStringDB>('updateFormFieldString'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormFieldStringService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormFieldStringService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
