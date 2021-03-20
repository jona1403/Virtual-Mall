import { Component, OnInit } from '@angular/core';
import { TiendasService } from "../../services/tiendas/tiendas.service";
import { DbTiendas } from "../../models/tiendas/db-tiendas";
import { Tienda } from "../../models/tienda/tienda";

@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css']
})
export class TiendasComponent implements OnInit {


  base: DbTiendas;
  lista_tiendas: Tienda[]=[];
  constructor(private tiendaService: TiendasService) {
    this.tiendaService.getTiendas().subscribe((db: any)=>{
      this.base = db;
    })
   }

  ngOnInit(): void {
  }

}
