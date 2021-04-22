import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Cl } from '../../models/cl/cl';
import { ClaveService } from '../../services/clave/clave.service';

@Component({
  selector: 'app-clave',
  templateUrl: './clave.component.html',
  styleUrls: ['./clave.component.css']
})
export class ClaveComponent implements OnInit {
  clave = new FormControl('')
  mostrarMensaje = false
  mostrarMensajeError = false
  constructor(private claveservice: ClaveService) { }

  ngOnInit(): void {
  }
  enviarclave(){
    const p: Cl={
      cla : this.clave.value
    }
    
    this.claveservice.postclave(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.clave.setValue("");
    }, (err) => {
      this.mostrarMensajeError = true;
    }
    )
  }
  desactivarMensaje(){
    this.mostrarMensaje = false;
    this.mostrarMensajeError = false;
  }
}
