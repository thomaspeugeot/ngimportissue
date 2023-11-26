import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NgimportissuespecificComponent } from './ngimportissuespecific.component';

describe('NgimportissuespecificComponent', () => {
  let component: NgimportissuespecificComponent;
  let fixture: ComponentFixture<NgimportissuespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [NgimportissuespecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(NgimportissuespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
