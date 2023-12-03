import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CountryDB } from './country-db'
import { CountryService } from './country.service'

import { HelloDB } from './hello-db'
import { HelloService } from './hello.service'

export const StackType = "ngimportissue/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  Countrys_array = new Array<CountryDB>() // array of repo instances
  Countrys = new Map<number, CountryDB>() // map of repo instances
  Countrys_batch = new Map<number, CountryDB>() // same but only in last GET (for finding repo instances to delete)

  Hellos_array = new Array<HelloDB>() // array of repo instances
  Hellos = new Map<number, HelloDB>() // map of repo instances
  Hellos_batch = new Map<number, HelloDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'Country':
        return this.Countrys_array as unknown as Array<Type>
      case 'Hello':
        return this.Hellos_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'Country':
        return this.Countrys as unknown as Map<number, Type>
      case 'Hello':
        return this.Hellos as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private countryService: CountryService,
    private helloService: HelloService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template 
    Observable<CountryDB[]>,
    Observable<HelloDB[]>,
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template
      this.countryService.getCountrys(this.GONG__StackPath, this.frontRepo),
      this.helloService.getHellos(this.GONG__StackPath, this.frontRepo),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template
      this.countryService.getCountrys(this.GONG__StackPath, this.frontRepo),
      this.helloService.getHellos(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            countrys_,
            hellos_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var countrys: CountryDB[]
            countrys = countrys_ as CountryDB[]
            var hellos: HelloDB[]
            hellos = hellos_ as HelloDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Countrys_array = countrys

            // clear the map that counts Country in the GET
            this.frontRepo.Countrys_batch.clear()

            countrys.forEach(
              country => {
                this.frontRepo.Countrys.set(country.ID, country)
                this.frontRepo.Countrys_batch.set(country.ID, country)
              }
            )

            // clear countrys that are absent from the batch
            this.frontRepo.Countrys.forEach(
              country => {
                if (this.frontRepo.Countrys_batch.get(country.ID) == undefined) {
                  this.frontRepo.Countrys.delete(country.ID)
                }
              }
            )

            // sort Countrys_array array
            this.frontRepo.Countrys_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Hellos_array = hellos

            // clear the map that counts Hello in the GET
            this.frontRepo.Hellos_batch.clear()

            hellos.forEach(
              hello => {
                this.frontRepo.Hellos.set(hello.ID, hello)
                this.frontRepo.Hellos_batch.set(hello.ID, hello)
              }
            )

            // clear hellos that are absent from the batch
            this.frontRepo.Hellos.forEach(
              hello => {
                if (this.frontRepo.Hellos_batch.get(hello.ID) == undefined) {
                  this.frontRepo.Hellos.delete(hello.ID)
                }
              }
            )

            // sort Hellos_array array
            this.frontRepo.Hellos_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            countrys.forEach(
              country => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Hello redeeming
                {
                  let _hello = this.frontRepo.Hellos.get(country.CountryPointersEncoding.HelloID.Int64)
                  if (_hello) {
                    country.Hello = _hello
                  }
                }
                // insertion point for pointers decoding
                country.AlternateHellos = new Array<HelloDB>()
                for (let _id of country.CountryPointersEncoding.AlternateHellos) {
                  let _hello = this.frontRepo.Hellos.get(_id)
                  if (_hello != undefined) {
                    country.AlternateHellos.push(_hello!)
                  }
                }
              }
            )
            hellos.forEach(
              hello => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // CountryPull performs a GET on Country of the stack and redeem association pointers 
  CountryPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.countryService.getCountrys(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            countrys,
          ]) => {
            // init the array
            this.frontRepo.Countrys_array = countrys

            // clear the map that counts Country in the GET
            this.frontRepo.Countrys_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            countrys.forEach(
              country => {
                this.frontRepo.Countrys.set(country.ID, country)
                this.frontRepo.Countrys_batch.set(country.ID, country)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Hello redeeming
                {
                  let _hello = this.frontRepo.Hellos.get(country.CountryPointersEncoding.HelloID.Int64)
                  if (_hello) {
                    country.Hello = _hello
                  }
                }
              }
            )

            // clear countrys that are absent from the GET
            this.frontRepo.Countrys.forEach(
              country => {
                if (this.frontRepo.Countrys_batch.get(country.ID) == undefined) {
                  this.frontRepo.Countrys.delete(country.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // HelloPull performs a GET on Hello of the stack and redeem association pointers 
  HelloPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.helloService.getHellos(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            hellos,
          ]) => {
            // init the array
            this.frontRepo.Hellos_array = hellos

            // clear the map that counts Hello in the GET
            this.frontRepo.Hellos_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            hellos.forEach(
              hello => {
                this.frontRepo.Hellos.set(hello.ID, hello)
                this.frontRepo.Hellos_batch.set(hello.ID, hello)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear hellos that are absent from the GET
            this.frontRepo.Hellos.forEach(
              hello => {
                if (this.frontRepo.Hellos_batch.get(hello.ID) == undefined) {
                  this.frontRepo.Hellos.delete(hello.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getCountryUniqueID(id: number): number {
  return 31 * id
}
export function getHelloUniqueID(id: number): number {
  return 37 * id
}
