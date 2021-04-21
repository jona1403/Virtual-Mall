import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarusuariosComponent } from './cargarusuarios.component';

describe('CargarusuariosComponent', () => {
  let component: CargarusuariosComponent;
  let fixture: ComponentFixture<CargarusuariosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarusuariosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarusuariosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
