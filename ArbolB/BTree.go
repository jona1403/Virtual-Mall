package ArbolB

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

var contador int

type DB_Users struct{
	Usuarios []User
}

type User struct{
	Dpi int
	Nombre string
	Correo string
	Password string
	Cuenta string
}

type Nodo struct{
	Hoja bool
	Cantidad int
	Claves [5]User
	Hijos [6]*Nodo
	Padre *Nodo

}

type ArbolB struct{
	Raiz *Nodo
	Grado int
	Enmedio int
}

//Metodos nodos

func (nodo *Nodo) insertar(usuario User){
	nodo.Claves[nodo.Cantidad] = usuario
	nodo.Cantidad++
	if nodo.Cantidad > 1{
		nodo.Claves = sort(nodo.Claves)
	}
}

func (nodo Nodo) PosNodo() int{
	for i:= 0; i<5; i++{
		if *nodo.Padre.Hijos[i] == nodo{
			return i
		}
	}
	return -1
}

func (nodo Nodo) buscar(usuario User, Usuarios [5]User)(bool){
	for i:= 0; i<5;i++{
		if usuario.Dpi == Usuarios[i].Dpi{
			return true
		}
	}
	return false
}

func sort(array [5]User)([5]User){
	var aux User
	var index int
	for i:= 0; i<5; i++{
		if array[i].Dpi == 0{
			index = i
			break
		}
	}
	if index == 0{
		index = 5
	}
	for i:= 0; i < index-1; i++{
		for j := i+1; j < index; j++{
			if array[i].Dpi > array[j].Dpi{
				aux = array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	return array
}

//Metodos arbol

func (arbol *ArbolB) Insertar(nuevo User){
	arbol.Raiz = arbol._insertar(nuevo, arbol.Raiz)
}

func recorrerArbol(nombrePadre string, hijo* Nodo, textoActual string) string{
	nombrehijo := "Nodo"+strconv.FormatInt(int64(contador), 10)
	contador++
	textoActual+= nombrehijo
	textoActual +=`[shape=none label=<`
	textoActual+=`<table cellspacing="0" border="0" cellborder="1">`
	textoActual+= "<tr>"
	for i:=0; i<5; i++{
		if hijo.Claves[i].Dpi != 0{
			textoActual += "<td>DPI: "+strconv.FormatInt(int64(hijo.Claves[i].Dpi), 10)+"<br/>"
			textoActual += "Nombre: "+hijo.Claves[i].Nombre+"<br/>"
			textoActual += "Correo: "+hijo.Claves[i].Correo+"<br/>"
			textoActual += "Password: "+hijo.Claves[i].Password+"<br/>"
			textoActual += "Cuenta: "+hijo.Claves[i].Cuenta+"<br/></td>"
		}else{
			break
		}
	}
	textoActual+= "</tr>"
	textoActual+="</table>"
	textoActual+=`
	>];
	`
	textoActual+=nombrePadre+"->"+nombrehijo+";\n"
	for i:= 0; i<6; i++{
		if hijo.Hijos[i] != nil{
			textoActual=recorrerArbol(nombrehijo, hijo.Hijos[i], textoActual)
		}else{
			break
		}
	}
	return textoActual
}

func CreateDot(nodo* Nodo) string{
	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=`Nodo0[shape=none label=<`
	grafo+=`<table cellspacing="0" border="0" cellborder="1">`
	grafo+= "<tr>"
	for i:=0; i<5; i++{
		if nodo.Claves[i].Dpi != 0{
			grafo += "<td>DPI: "+strconv.FormatInt(int64(nodo.Claves[i].Dpi), 10)+"<br/>"
			grafo += "Nombre: "+nodo.Claves[i].Nombre+"<br/>"
			grafo += "Correo: "+nodo.Claves[i].Correo+"<br/>"
			grafo += "Password: "+nodo.Claves[i].Password+"<br/>"
			grafo += "Cuenta: "+nodo.Claves[i].Cuenta+"<br/></td>"
		}else{
			break
		}
	}
	grafo+= "</tr>"
	grafo+="</table>"
	grafo+=`
	>];
	`
	contador = 1
	for i:= 0; i<6; i++{
		if nodo.Hijos[i] != nil{
			grafo=recorrerArbol("Nodo0", nodo.Hijos[i], grafo)
		}else{
			break
		}
	}

	grafo+="}"
	return grafo
}

func encriptar(datos []byte, conn string) string{
	h:= hmac.New(sha256.New, []byte(conn))
	h.Write(datos)
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func recorrerArbolcifrado(nombrePadre string, hijo* Nodo, textoActual string, password string, tipo int) string{
	nombrehijo := "Nodo"+strconv.FormatInt(int64(contador), 10)
	contador++
	textoActual+= nombrehijo
	textoActual +=`[shape=none label=<`
	textoActual+=`<table cellspacing="0" border="0" cellborder="1">`
	textoActual+= "<tr>"
	for i:=0; i<5; i++{
		if hijo.Claves[i].Dpi != 0{
			texto := encriptar([]byte(strconv.FormatInt(int64(hijo.Claves[i].Dpi), 10)), password)
			textoActual += "<td>DPI: "+texto+"<br/>"

			texto = encriptar([]byte(hijo.Claves[i].Correo), password)
			textoActual += "Correo: "+texto+"<br/>"
			texto = encriptar([]byte(hijo.Claves[i].Password), password)
			textoActual += "Password: "+texto+"<br/>"

			if tipo == 1{
				texto = encriptar([]byte(hijo.Claves[i].Nombre), password)
				textoActual += "Nombre: "+texto+"<br/>"
				texto = encriptar([]byte(hijo.Claves[i].Cuenta), password)
				textoActual += "Cuenta: "+texto+"<br/></td>"
			}else{
				textoActual += "Nombre: "+hijo.Claves[i].Nombre+"<br/>"
				textoActual += "Cuenta: "+hijo.Claves[i].Cuenta+"<br/></td>"
			}
		}else{
			break
		}
	}
	textoActual+= "</tr>"
	textoActual+="</table>"
	textoActual+=`
	>];
	`
	textoActual+=nombrePadre+"->"+nombrehijo+";\n"
	for i:= 0; i<6; i++{
		if hijo.Hijos[i] != nil{
			textoActual=recorrerArbolcifrado(nombrehijo, hijo.Hijos[i], textoActual, password, tipo)
		}else{
			break
		}
	}
	return textoActual
}

func CreateDotcifrado(nodo *Nodo, password string, tipo int) string{
	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=`Nodo0[shape=none label=<`
	grafo+=`<table cellspacing="0" border="0" cellborder="1">`
	grafo+= "<tr>"
	for i:=0; i<5; i++{
		if nodo.Claves[i].Dpi != 0{
			texto := encriptar([]byte(strconv.FormatInt(int64(nodo.Claves[i].Dpi), 10)), password)
			grafo += "<td>DPI: "+texto+"<br/>"

			texto = encriptar([]byte(nodo.Claves[i].Correo), password)
			grafo += "Correo: "+texto+"<br/>"
			texto = encriptar([]byte(nodo.Claves[i].Password), password)
			grafo += "Password: "+texto+"<br/>"
			if tipo == 1{
				texto = encriptar([]byte(nodo.Claves[i].Nombre), password)
				grafo += "Nombre: "+texto+"<br/>"
				texto = encriptar([]byte(nodo.Claves[i].Cuenta), password)
				grafo += "Cuenta: "+texto+"<br/></td>"
			}else{
				grafo += "Nombre: "+nodo.Claves[i].Nombre+"<br/>"
				grafo += "Cuenta: "+nodo.Claves[i].Cuenta+"<br/></td>"
			}
		}else{
			break
		}
	}
	grafo+= "</tr>"
	grafo+="</table>"
	grafo+=`
	>];
	`
	contador = 1
	for i:= 0; i<6; i++{
		if nodo.Hijos[i] != nil{
			grafo=recorrerArbolcifrado("Nodo0", nodo.Hijos[i], grafo, password, tipo)
		}else{
			break
		}
	}

	grafo+="}"
	return grafo
}


func (arbol ArbolB) _insertar(nuevo User, temp *Nodo) (*Nodo){
	if temp.Hijos[0] == nil{
		temp.insertar(nuevo)
	}else{
		encontrado := false
		for i:= 0; i<temp.Cantidad; i++{
			if nuevo.Dpi < temp.Claves[i].Dpi{
				encontrado = true
				arbol._insertar(nuevo, temp.Hijos[i])
				break
			}
		}
		if(!encontrado){
			arbol._insertar(nuevo, temp.Hijos[temp.Cantidad])
		}
	}
	if(temp.Claves[4].Dpi != 0) {
		if temp.Padre == nil {
			c := temp
			temp = &Nodo{Padre: nil, Hijos: [6]*Nodo{}, Claves: [5]User{}, Hoja: true, Cantidad: 0}
			temp.insertar(c.Claves[2])
			temp.Hijos[0] = &Nodo{Padre: temp, Hoja: true}
			temp.Hijos[1] = &Nodo{Padre: temp, Hoja: true}
			for i:= 0; i< 2; i++{
				temp.Hijos[0].insertar(c.Claves[i])
			}
			for i:= 3; i< 5; i++{
				temp.Hijos[1].insertar(c.Claves[i])
			}
			temp.Hoja = false
			tienehijos:= true
			if c.Hijos[0] == nil{
				tienehijos = false
			}
			if tienehijos{
				for i:= 0; i<3; i++{
					temp.Hijos[0].Hijos[i] = c.Hijos[i]
					temp.Hijos[0].Hijos[i].Padre = temp.Hijos[0]
				}
				for i:= 3; i<6; i++{
					temp.Hijos[1].Hijos[i-3] = c.Hijos[i]
					temp.Hijos[1].Hijos[i-3].Padre = temp.Hijos[1]
				}
			}
		} else {
			claveMedia := temp.Claves[2]
			temp.Padre.insertar(claveMedia)
			tieneHijos := true
			for i:=0; i< 6;i++{
				if temp.Hijos[i] == nil{
					tieneHijos = false
					break
				}
			}
			var index int
			for index = 0;index< temp.Padre.Cantidad; index++{
				if temp.Padre.Claves[index] == claveMedia{
					break
				}
			}
			for i := temp.Padre.Cantidad; i> index+1; i--{
				temp.Padre.Hijos[i] = temp.Padre.Hijos[i-1]
			}
			aux := temp
			temp.Padre.Hijos[index] = &Nodo{Padre: temp.Padre}
			for i:= 0; i < 2; i++{
				temp.Padre.Hijos[index].insertar(aux.Claves[i])
			}
			temp.Padre.Hijos[index+1] = &Nodo{Padre: temp.Padre, Hoja: true}
			for i:= 3;i<5; i++{
				temp.Padre.Hijos[index+1].insertar(aux.Claves[i])
			}
			if tieneHijos{
				for i:= 0; i<3; i++{
					temp.Padre.Hijos[index].Hijos[i] = aux.Hijos[i]
					temp.Padre.Hijos[index].Hijos[i].Padre = temp.Padre.Hijos[index]
				}
				for i:= 3; i<6; i++{
					temp.Padre.Hijos[index+1].Hijos[i-3] = aux.Hijos[i]
					temp.Padre.Hijos[index+1].Hijos[i-3].Padre = temp.Padre.Hijos[index+1]
				}
			}
		}
	}
	return temp
}

func Insercionmasiva(db DB_Users, Ab ArbolB) ArbolB{
	for i:= 0; i< len(db.Usuarios); i++{
		Ab.Insertar(db.Usuarios[i])
	}
	return Ab
}