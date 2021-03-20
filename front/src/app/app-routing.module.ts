  
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { CargarpedidosComponent } from "./componentes/cargarpedidos/cargarpedidos.component";
import { VerpedidoComponent } from "./componentes/verpedido/verpedido.component";
import { CargarproductosComponent } from './componentes/cargarproductos/cargarproductos.component';
const routes: Routes = [
    {
        path: '',
        component:InicioComponent,
    },
    {
        path: 'cargartiendas',
        component:CargarTiendasComponent,
    },
    {
        path: 'tiendas',
        component:TiendasComponent,
    },
    {
        path: 'cargarproductos',
        component:CargarproductosComponent,
    },
    {
        path: 'cargarpedidos',
        component:CargarpedidosComponent,
    },
    {
        path: 'pedido',
        component:VerpedidoComponent,
    },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }