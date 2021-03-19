import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-cargarproductos',
  templateUrl: './cargarproductos.component.html',
  styleUrls: ['./cargarproductos.component.css']
})
export class CargarproductosComponent implements OnInit {

  productos = new FormControl('')
  
  mostrarMensaje = false
  mostrarMensajeError = false

  constructor() { }

  ngOnInit(): void {
  }

}
