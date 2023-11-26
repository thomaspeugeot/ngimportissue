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

import { LineDB } from './line-db';
import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateDB } from './animate-db'

@Injectable({
  providedIn: 'root'
})
export class LineService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LineServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private linesUrl: string

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
    this.linesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/lines';
  }

  /** GET lines from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB[]> {
    return this.getLines(GONG__StackPath, frontRepo)
  }
  getLines(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<LineDB[]>(this.linesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<LineDB[]>('getLines', []))
      );
  }

  /** GET line by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {
    return this.getLine(id, GONG__StackPath, frontRepo)
  }
  getLine(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.linesUrl}/${id}`;
    return this.http.get<LineDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched line id=${id}`)),
      catchError(this.handleError<LineDB>(`getLine id=${id}`))
    );
  }

  /** POST: add a new line to the server */
  post(linedb: LineDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {
    return this.postLine(linedb, GONG__StackPath, frontRepo)
  }
  postLine(linedb: LineDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    linedb.LinePointersEncoding.Animates = []
    for (let _animate of linedb.Animates) {
      linedb.LinePointersEncoding.Animates.push(_animate.ID)
    }
    linedb.Animates = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LineDB>(this.linesUrl, linedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linedb.Animates = new Array<AnimateDB>()
        for (let _id of linedb.LinePointersEncoding.Animates) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            linedb.Animates.push(_animate!)
          }
        }
        // this.log(`posted linedb id=${linedb.ID}`)
      }),
      catchError(this.handleError<LineDB>('postLine'))
    );
  }

  /** DELETE: delete the linedb from the server */
  delete(linedb: LineDB | number, GONG__StackPath: string): Observable<LineDB> {
    return this.deleteLine(linedb, GONG__StackPath)
  }
  deleteLine(linedb: LineDB | number, GONG__StackPath: string): Observable<LineDB> {
    const id = typeof linedb === 'number' ? linedb : linedb.ID;
    const url = `${this.linesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<LineDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted linedb id=${id}`)),
      catchError(this.handleError<LineDB>('deleteLine'))
    );
  }

  /** PUT: update the linedb on the server */
  update(linedb: LineDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {
    return this.updateLine(linedb, GONG__StackPath, frontRepo)
  }
  updateLine(linedb: LineDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LineDB> {
    const id = typeof linedb === 'number' ? linedb : linedb.ID;
    const url = `${this.linesUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    linedb.LinePointersEncoding.Animates = []
    for (let _animate of linedb.Animates) {
      linedb.LinePointersEncoding.Animates.push(_animate.ID)
    }
    linedb.Animates = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<LineDB>(url, linedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        linedb.Animates = new Array<AnimateDB>()
        for (let _id of linedb.LinePointersEncoding.Animates) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            linedb.Animates.push(_animate!)
          }
        }
        // this.log(`updated linedb id=${linedb.ID}`)
      }),
      catchError(this.handleError<LineDB>('updateLine'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in LineService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("LineService" + error); // log to console instead

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