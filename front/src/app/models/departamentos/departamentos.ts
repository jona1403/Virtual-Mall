import { Tienda } from "../../models/tienda/tienda"

export class Departamentos {
    Nombre: string
    Tiendas: Tienda[]
    constructor(_Nombre: string, _Tiendas: Tienda[]){
        this.Nombre= _Nombre
        this.Tiendas= _Tiendas 
    }
}
