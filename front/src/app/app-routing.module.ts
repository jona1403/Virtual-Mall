  
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
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
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }