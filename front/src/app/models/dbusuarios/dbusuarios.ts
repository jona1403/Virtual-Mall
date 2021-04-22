import { Usuario } from "../../models/usuario/usuario";
export class Dbusuarios {
    Usuarios: Usuario[]

    constructor(_Usuarios: Usuario[]){
        this.Usuarios = _Usuarios
    }
}
