import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { PedidosService } from "../../services/pedidos/pedidos.service";

@Component({
  selector: 'app-cargarpedidos',
  templateUrl: './cargarpedidos.component.html',
  styleUrls: ['./cargarpedidos.component.css']
})
export class CargarpedidosComponent implements OnInit {

  pedidos = new FormControl('')
  mostrarMensaje = false
  mostrarMensajeError = false

  constructor(private pedidosservice: PedidosService) { }

  ngOnInit(): void {
  }

  cargarpedidos(){
    const p: JSON =  JSON.parse(this.pedidos.value);
    //const p: Bdinventario = new Bdinventario(this.productos.value);
    this.pedidosservice.postPedidos(p).subscribe((res: any) =>{
      this.mostrarMensaje = true;
      this.pedidos.setValue("");
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
