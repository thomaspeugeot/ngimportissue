# esbuild import path resolution issue with angular v17 for modules outsides the workspace

```bash
git clone https://github.com/thomaspeugeot/ngimportissue.git
```

## Situation

A workspace has been created with angular version 17.

```bash
% ng version
...
Angular CLI: 17.0.3
```

the workspace has been created with the following config

```bash
ng new ng --defaults=true --minimal=true --no-standalone --routing --ssr=false
```

## The issue arises when ng tries to compiles modules outsides the workspace.

As seen in [ng/src/app/app.module.ts](https://github.com/thomaspeugeot/ngimportissue/blob/8a11d9dbe0a8a2b233f0e2073cc67723e63fb9a0/ng/src/app/app.module.ts), the application imports some modules that are outside the workspace. One of them is `GongModule`. 

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

As seen in [vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/lib/gong.module.ts](https://github.com/thomaspeugeot/ngimportissue/blob/8a11d9dbe0a8a2b233f0e2073cc67723e63fb9a0/vendor/github.com/fullstack-lang/gong/ng/projects/gong/src/lib/gong.module.ts), `GongModule` itself imports angular  modules


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

### This is new with ng 17

This problem did not occur with ng v16

### This is an issue related to the import path

going verbose and having a trace in [ng/build_trace.txt](https://github.com/thomaspeugeot/ngimportissue/blob/8a11d9dbe0a8a2b233f0e2073cc67723e63fb9a0/ng/build_trace.txt), one better understands what's wrong

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

It seems the builder is mislead by the the import path. It does not search the `esm2022` or `fesm2022` where the index file is present. Indeed `ng/node_modules/@angular/common/esm2022/http/http.mjs` is present (for information, this file is present but not in the git repo also, you need to perfom `npm i` to have it present).

The import path in [ng/tsconfig.json](https://github.com/thomaspeugeot/ngimportissue/blob/8a11d9dbe0a8a2b233f0e2073cc67723e63fb9a0/ng/tsconfig.json) that worked with ng v16 is configured to work with import path outside the workspace.

```json
    "paths": {
      "*": [
        "./node_modules/*"
      ],
```

It does not work with ng 17.

### There is a Workaround

If one configures the project as a `browser` instead of `application`, the compilation works.
 It is this configuration that is created when one migrates the projects from ng 16 to ng 17.

Below is the diff for [ng/angular.json](https://github.com/thomaspeugeot/ngimportissue/blob/8a11d9dbe0a8a2b233f0e2073cc67723e63fb9a0/ng/angular.json)

```
--- a/ng/angular.json
+++ b/ng/angular.json
@@ -5,7 +5,7 @@
     "ng": {
       "architect": {
         "build": {
-          "builder": "@angular-devkit/build-angular:application",
+          "builder": "@angular-devkit/build-angular:browser",
           "configurations": {
             "development": {
               "extractLicenses": false,
@@ -34,7 +34,7 @@
               "src/favicon.ico",
               "src/assets"
             ],
-            "browser": "src/main.ts",
+            "main": "src/main.ts",
             "index": "src/index.html",
:...skipping...
diff --git a/ng/angular.json b/ng/angular.json
index 288f403..f4e5a7c 100644
--- a/ng/angular.json
+++ b/ng/angular.json
@@ -5,7 +5,7 @@
     "ng": {
       "architect": {
         "build": {
-          "builder": "@angular-devkit/build-angular:application",
+          "builder": "@angular-devkit/build-angular:browser",
           "configurations": {
             "development": {
               "extractLicenses": false,
@@ -34,7 +34,7 @@
               "src/favicon.ico",
               "src/assets"
             ],
-            "browser": "src/main.ts",
+            "main": "src/main.ts",
             "index": "src/index.html",
             "outputPath": "dist/ng",
             "polyfills": [
```

## Questions

I would like to undestand why the compilation fails.

I am looking for the proper tsconfig.js config to have the compilation works. Other solutions could be suitable of course.


