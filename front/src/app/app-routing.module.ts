  
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { CargarusuariosComponent } from './componentes/cargarusuarios/cargarusuarios.component';
import {CrearUsuarioComponent} from './componentes/crear-usuario/crear-usuario.component';
import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { CargarpedidosComponent } from "./componentes/cargarpedidos/cargarpedidos.component";
import { VerpedidoComponent } from "./componentes/verpedido/verpedido.component";
import { CargarproductosComponent } from './componentes/cargarproductos/cargarproductos.component';
import { ClaveComponent } from './componentes/clave/clave.component';
const routes: Routes = [
    {
        path: 'clave',
        component:ClaveComponent,
    },
    {
        path: 'cargarusuarios',
        component:CargarusuariosComponent,
    },
    {
        path: 'crearusuario',
        component:CrearUsuarioComponent,
    },
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