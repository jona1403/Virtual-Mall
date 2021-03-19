import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DbTiendas } from 'src/app/models/tiendas/db-tiendas';
import { TiendasService } from "../../services/tiendas/tiendas.service";

@Component({
  selector: 'app-cargar-tiendas',
  templateUrl: './cargar-tiendas.component.html',
  styleUrls: ['./cargar-tiendas.component.css']
})
export class CargarTiendasComponent implements OnInit {

  tiendas = new FormControl('')
  
  mostrarMensaje = false
  mostrarMensajeError = false

  constructor(private tiendasService: TiendasService) {  }

  agregarTiendas(){
    const t: DbTiendas = new DbTiendas(this.tiendas.value);
    this.tiendasService.postTienda(t).subscribe((res: any) => {
      this.mostrarMensaje = true;
      this.tiendas.setValue("");
    }, (err)=>{
      this.mostrarMensajeError = true;
    })
  }

  desactivarMensaje(){
    this.mostrarMensaje = false;
    this.mostrarMensajeError = false;
  }

  ngOnInit(): void {
  }

}
