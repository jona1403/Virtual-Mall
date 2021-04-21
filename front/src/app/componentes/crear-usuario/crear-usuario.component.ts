import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-crear-usuario',
  templateUrl: './crear-usuario.component.html',
  styleUrls: ['./crear-usuario.component.css']
})
export class CrearUsuarioComponent implements OnInit {
  dpi = new FormControl('')
  nombre = new FormControl('')
  correo = new FormControl('')
  password = new FormControl('')
  cuenta = new FormControl('')
  mostrarMensaje = false
  mostrarMensajeError = false
  constructor() { }

  ngOnInit(): void {
  }
  crearusuario(){
    /*const p: JSON =  JSON.parse(this.usuarios.value);
    //const p: Bdinventario = new Bdinventario(this.productos.value);
    this.pedidosservice.postPedidos(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.usuarios.setValue("");
    }, (err) => {
      this.mostrarMensajeError = true;
    }
    )*/
  }

  desactivarMensaje(){
    this.mostrarMensaje = false;
    this.mostrarMensajeError = false;
  }
}
