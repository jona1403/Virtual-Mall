import { TestBed } from '@angular/core/testing';

import { ClaveService } from './clave.service';

describe('ClaveService', () => {
  let service: ClaveService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClaveService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
