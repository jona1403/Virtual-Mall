import { Departamentos } from "../../models/departamentos/departamentos"
export class Datosindices {
    Indice: string
    Departamentos: Departamentos[]
    
    constructor(_Indice: string, _Departamentos: Departamentos[]){
        this.Indice = _Indice
        this.Departamentos = _Departamentos
    }
}
