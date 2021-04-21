import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { UsuariosService } from "../../services/usuarios/usuarios.service";

@Component({
  selector: 'app-cargarusuarios',
  templateUrl: './cargarusuarios.component.html',
  styleUrls: ['./cargarusuarios.component.css']
})
export class CargarusuariosComponent implements OnInit {

  usuarios = new FormControl('')
  mostrarMensaje = false
  mostrarMensajeError = false
  constructor(private usuariosservice: UsuariosService) { }

  ngOnInit(): void {
  }

  cargarusuarios(){
    const p: JSON =  JSON.parse(this.usuarios.value);
    //const p: Bdinventario = new Bdinventario(this.productos.value);
    this.usuariosservice.postUsuarios(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.usuarios.setValue("");
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
