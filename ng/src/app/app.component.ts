import { Component, OnInit } from '@angular/core';

import * as ngimportissue from 'ngimportissue'

import * as gongtable from 'gongtable'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Ngimportissue Data/Model'
  view = this.default

  views: string[] = [this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "ngimportissue"
  StackType = ngimportissue.StackType

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
