import { Datosindices } from "../../models/datosindices/datosindices"

export class DbTiendas {
    Datos: Datosindices[]

    constructor(_Datos: Datosindices[]){
        this.Datos = _Datos
    }
}
