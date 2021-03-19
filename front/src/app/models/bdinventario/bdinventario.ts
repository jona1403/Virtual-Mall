import { Inventariotienda } from "../../models/inventariotienda/inventariotienda";
export class Bdinventario {
    Inventarios: Inventariotienda[]

    constructor(_Inventarios: Inventariotienda[]){
        this.Inventarios = _Inventarios
    }
}
