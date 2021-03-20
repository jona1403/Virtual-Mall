import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarpedidosComponent } from './cargarpedidos.component';

describe('CargarpedidosComponent', () => {
  let component: CargarpedidosComponent;
  let fixture: ComponentFixture<CargarpedidosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarpedidosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarpedidosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
