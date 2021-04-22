import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule, FormsModule } from "@angular/forms";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from "@angular/common/http"

import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { CargarproductosComponent } from './componentes/cargarproductos/cargarproductos.component';
import { VerproductosComponent } from './componentes/verproductos/verproductos.component';
import { CargarpedidosComponent } from './componentes/cargarpedidos/cargarpedidos.component';
import { VerpedidoComponent } from './componentes/verpedido/verpedido.component';
import { CrearUsuarioComponent } from './componentes/crear-usuario/crear-usuario.component';
import { CargarusuariosComponent } from './componentes/cargarusuarios/cargarusuarios.component';
import { ClaveComponent } from './componentes/clave/clave.component';

@NgModule({
  declarations: [
    AppComponent,
    TiendasComponent,
    CargarTiendasComponent,
    InicioComponent,
    CargarproductosComponent,
    VerproductosComponent,
    CargarpedidosComponent,
    VerpedidoComponent,
    CrearUsuarioComponent,
    CargarusuariosComponent,
    ClaveComponent
  ],
  imports: [
    BrowserModule, 
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
