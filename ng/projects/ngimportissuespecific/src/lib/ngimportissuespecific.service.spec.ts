import { TestBed } from '@angular/core/testing';

import { NgimportissuespecificService } from './ngimportissuespecific.service';

describe('NgimportissuespecificService', () => {
  let service: NgimportissuespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NgimportissuespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
