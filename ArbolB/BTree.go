package ArbolB

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
	Hijos [5]*Nodo
	Padre *Nodo

}

type ArbolB struct{
	Raiz *Nodo
	Grado int
	Enmedio int
}

//Funciones de los nodos o paginas
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
	for i:= 0; i < 5-1; i++{
		for j := i+1; j < 5; j++{
			if array[i].Dpi > array[j].Dpi{
				aux = array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	return array
}


//funciones del arbol


func (arbol *ArbolB) insertar(nuevo User){
arbol.Raiz = arbol._insertar(nuevo, arbol.Raiz)
}

func (arbol ArbolB) _insertar(nuevo User, temp *Nodo) (*Nodo){
	if temp.Hoja{
		temp.insertar(nuevo)
	}else{
		encontrado := false
		for i:= 0; i<temp.Cantidad-1; i++{
			if nuevo.Dpi < temp.Claves[i].Dpi{
				encontrado = true
				temp.Hijos[i] = arbol._insertar(nuevo, temp.Hijos[i])
				break
			}
		}
		if(!encontrado){
			temp.Hijos[temp.Cantidad] = arbol._insertar(nuevo, temp.Hijos[temp.Cantidad])
		}
	}
	if(temp.Cantidad == 5) {
		if temp.Padre == nil {
			c := temp
			temp = &Nodo{}
			temp.insertar(c.Claves[2])
			temp.Hijos[0] = &Nodo{Padre: temp}
			temp.Hijos[1] = &Nodo{Padre: temp}
			for i:= 0; i< 2; i++{
				temp.Hijos[0].insertar(c.Claves[i])
			}
			for i:= 3; i< 5; i++{
				temp.Hijos[1].insertar(c.Claves[i])
			}
			temp.Hoja = false
		} else {
			claveMedia := temp.Claves[2]
			temp.Padre.insertar(claveMedia)
			var index int
			for index = 0;index< temp.Padre.Cantidad; index++{
				if temp.Padre.Claves[index] == claveMedia{
					break
				}
			}
			for i := temp.Padre.Cantidad; i> index+1; i--{
				temp.Padre.Hijos[i] = temp.Padre.Hijos[i-1]
			}
			//Falta codigo
			aux := temp
			temp.Padre.Hijos[index] = &Nodo{Padre: temp.Padre}
			for i:= 0; i < 2; i++{
				temp.Padre.Hijos[index].insertar(aux.Claves[i])
			}

		}
	}
	return temp
}