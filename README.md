# ng import path issue

## Sitation

workspace created with angular version 17

````
ng version

     _                      _                 ____ _     ___
    / \   _ __   __ _ _   _| | __ _ _ __     / ___| |   |_ _|
   / △ \ | '_ \ / _` | | | | |/ _` | '__|   | |   | |    | |
  / ___ \| | | | (_| | |_| | | (_| | |      | |___| |___ | |
 /_/   \_\_| |_|\__, |\__,_|_|\__,_|_|       \____|_____|___|
                |___/
    

Angular CLI: 17.0.3
Node: 20.10.0
Package Manager: npm 10.2.3
OS: darwin x64

Angular: 
... 

Package                      Version
------------------------------------------------------
@angular-devkit/architect    0.1700.3 (cli-only)
@angular-devkit/core         17.0.3 (cli-only)
@angular-devkit/schematics   17.0.3 (cli-only)
@schematics/angular          17.0.3 (cli-only)

ng new ng --defaults=true --minimal=true --no-standalone --routing --ssr=false
```

## Import path for the project

The application imports some modules that are outside the workspace. One of them is `GongModule``

ng/src/app/app.module.ts

```ts
...
import { GongModule } from 'gong'


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    ...
    GongModule,
```

`GongModule` itself imports angular  modules

vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/lib/gong.module.ts

```ts
import { NgModule } from '@angular/core';

@NgModule({
	declarations: [
	],
	imports: [
		CommonModule,

        ...
...
```

Angular cannot build the application

```
 % ng build
⠋ Building...✘ [ERROR] Could not resolve "@angular/common/http"

    ../vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/lib/commitnbfromback.service.ts:2:27:
      2 │ import { HttpParams } from '@angular/common/http';
        ╵                            ~~~~~~~~~~~~~~~~~~~~~~
...
```

## Preliminary analysis

This problem did not occur with ng v16

going verbose illuminates what's wrong

> % ng build --verbose > build_trace.txt 2>&1 

ng/build_trace.txt

```
Searching for "@angular/common/http" in "node_modules" directories starting from "/private/tmp/ngimportissue/vendor/github.com/fullstack-lang/gongdoc/ng/projects/gongdoc/src/lib"
    Matching "@angular/common/http" against "paths" in "/private/tmp/ngimportissue/ng/tsconfig.app.json"
      Using "/private/tmp/ngimportissue/ng" as "baseURL"
      Found a fuzzy match for "*" in "paths"
      Attempting to load "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http" as a file
        Checking for file "http"
        Checking for file "http.mjs"
        Checking for file "http.js"
        Checking for file "http.ts"
        Checking for file "http.tsx"
        Failed to find file "http"
      Attempting to load "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http" as a directory
        Read 2 entries for directory "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http"
        No "browser" map found in directory "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http"
        Failed to find file "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http/index.mjs"
        Failed to find file "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http/index.js"
        Failed to find file "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http/index.ts"
        Failed to find file "/private/tmp/ngimportissue/ng/node_modules/@angular/common/http/index.tsx"
    Parsed package name "@angular/common" and package subpath "./http"

✘ [ERROR] Could not resolve "@angular/common/http"
```

It seems the builder is mislead by the the import path. It does not search the `esm2022` or `fesm2022` where the index file is present. Indeed `ng/node_modules/@angular/common/esm2022/http/http.mjs` is present.

The import path that worked with ng v16 is configured to work with import path outside the workspace.

ng/tsconfig.json

```json
    "paths": {
      "*": [
        "./node_modules/*"
      ],
```

The does not work.



## Issue



