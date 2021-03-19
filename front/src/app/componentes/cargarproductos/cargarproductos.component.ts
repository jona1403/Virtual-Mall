import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Bdinventario } from 'src/app/models/bdinventario/bdinventario';
import { ProductosService } from "../../services/productos/productos.service";

@Component({
  selector: 'app-cargarproductos',
  templateUrl: './cargarproductos.component.html',
  styleUrls: ['./cargarproductos.component.css']
})
export class CargarproductosComponent implements OnInit {

  productos = new FormControl('')
  
  mostrarMensaje = false
  mostrarMensajeError = false

  constructor(private productosservice: ProductosService) { }

  ngOnInit(): void {
  }

  insertarProductos(){
    const p: Bdinventario = new Bdinventario(this.productos.value);
    this.productosservice.postProductos(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.productos.setValue("");
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
