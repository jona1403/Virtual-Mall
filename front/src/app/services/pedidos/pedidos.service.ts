import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from 'C:/Users/HUGO/Desktop/5to Semestre/Estructura de datos/VirtualMall/front/src/app/apiURL/baseURL';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PedidosService {

  
  constructor(private http: HttpClient) { }

  postPedidos(Pedidos): Observable<any>{
    const httpOptions = {
      headers : new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
    return this.http.post<any>(baseURL+'setPedidos', Pedidos, httpOptions)
  }

}
