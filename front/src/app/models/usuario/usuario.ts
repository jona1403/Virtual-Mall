export class Usuario {
    Dpi: number
    Nombre: string
    Correo: string
    Password: string
    Cuenta: number
    constructor(_Dpi:number,_Nombre:string, _Correo: string, _Password: string, _Cuenta: number){
        this.Dpi = _Dpi
        this.Nombre = _Nombre
        this.Correo = _Correo
        this.Password = _Password
        this.Cuenta = _Cuenta
    }
}
