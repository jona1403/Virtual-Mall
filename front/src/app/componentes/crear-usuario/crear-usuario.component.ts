import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Usuario } from '../../models/usuario/usuario';
import { UsuarioService } from '../../services/usuario/usuario.service';

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
  constructor(private usuarioservicio: UsuarioService) { }

  ngOnInit(): void {
  }
  crearusuario(){
    const p: Usuario={
      Dpi : Number(this.dpi.value),
      Nombre : this.nombre.value,
      Correo : this.correo.value,
      Password : this.password.value,
      Cuenta : this.cuenta.value
    }
    
    this.usuarioservicio.postUsuario(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.dpi.setValue("");
      this.nombre.setValue("");
      this.correo.setValue("");
      this.password.setValue("");
      this.cuenta.setValue("");
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
