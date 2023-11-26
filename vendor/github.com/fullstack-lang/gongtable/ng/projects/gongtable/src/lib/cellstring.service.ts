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

import { CellStringDB } from './cellstring-db';
import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class CellStringService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CellStringServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private cellstringsUrl: string

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
    this.cellstringsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/cellstrings';
  }

  /** GET cellstrings from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB[]> {
    return this.getCellStrings(GONG__StackPath, frontRepo)
  }
  getCellStrings(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CellStringDB[]>(this.cellstringsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CellStringDB[]>('getCellStrings', []))
      );
  }

  /** GET cellstring by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {
    return this.getCellString(id, GONG__StackPath, frontRepo)
  }
  getCellString(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.cellstringsUrl}/${id}`;
    return this.http.get<CellStringDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched cellstring id=${id}`)),
      catchError(this.handleError<CellStringDB>(`getCellString id=${id}`))
    );
  }

  /** POST: add a new cellstring to the server */
  post(cellstringdb: CellStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {
    return this.postCellString(cellstringdb, GONG__StackPath, frontRepo)
  }
  postCellString(cellstringdb: CellStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellStringDB>(this.cellstringsUrl, cellstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted cellstringdb id=${cellstringdb.ID}`)
      }),
      catchError(this.handleError<CellStringDB>('postCellString'))
    );
  }

  /** DELETE: delete the cellstringdb from the server */
  delete(cellstringdb: CellStringDB | number, GONG__StackPath: string): Observable<CellStringDB> {
    return this.deleteCellString(cellstringdb, GONG__StackPath)
  }
  deleteCellString(cellstringdb: CellStringDB | number, GONG__StackPath: string): Observable<CellStringDB> {
    const id = typeof cellstringdb === 'number' ? cellstringdb : cellstringdb.ID;
    const url = `${this.cellstringsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CellStringDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted cellstringdb id=${id}`)),
      catchError(this.handleError<CellStringDB>('deleteCellString'))
    );
  }

  /** PUT: update the cellstringdb on the server */
  update(cellstringdb: CellStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {
    return this.updateCellString(cellstringdb, GONG__StackPath, frontRepo)
  }
  updateCellString(cellstringdb: CellStringDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellStringDB> {
    const id = typeof cellstringdb === 'number' ? cellstringdb : cellstringdb.ID;
    const url = `${this.cellstringsUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CellStringDB>(url, cellstringdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated cellstringdb id=${cellstringdb.ID}`)
      }),
      catchError(this.handleError<CellStringDB>('updateCellString'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CellStringService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CellStringService" + error); // log to console instead

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
