import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from 'C:/Users/HUGO/Desktop/5to Semestre/Estructura de datos/VirtualMall/front/src/app/apiURL/baseURL';
import { Observable } from 'rxjs';
import { DbTiendas } from "../../models/tiendas/db-tiendas";

@Injectable({
  providedIn: 'root'
})
export class TiendasService {

  constructor(private http: HttpClient) {}

  postTienda(Tiendas): Observable<any>{
    const httpOptions = {
      headers : new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
    return this.http.post<any>(baseURL+'cargartienda', Tiendas, httpOptions)
  }


  getTiendas(): Observable<any>{
    const httpOptions = {
      headers : new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    }
    return this.http.get<any>(baseURL+'Tiendas', httpOptions)
  }
}
