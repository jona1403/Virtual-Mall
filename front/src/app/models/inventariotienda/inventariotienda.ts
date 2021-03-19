import { Producto } from "../../models/producto/producto";

export class Inventariotienda {
    Tienda : string
    Departamento: string
    Calificacion: number
    Productos: Producto[]

    constructor(_Tienda: string, _Departamento: string, _Calificacion: number, _Productos: Producto[]){
        this.Tienda = _Tienda
        this.Departamento = _Departamento
        this.Calificacion = _Calificacion
        this.Productos = _Productos
    }
}
