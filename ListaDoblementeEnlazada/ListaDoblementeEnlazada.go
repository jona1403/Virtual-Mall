package ListaDoblementeEnlazada

import "fmt"

//Tipo tienda a eliminar
type TiendaEliminar struct{
	Nombre string
	Categoria string
	Calificacion int
}

//Tipo ingresado para la busqueda
type TiendaIntroducida struct{
	Departamento string
	Nombre string
	Calificacion int
}

//Tipo tienda, tiene todos los atributos de las tiendas
type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
	Logo string
}

//Nodo, es utilizado para la lista doblemente enlazada
type Nodo struct {
	Anterior, Siguiente *Nodo
	Dato                Tienda
}

//Lista doblemente enlazada, contiene todas las tiendas que se encuentren dentro
//de la misma sección y mismo puntaje
type ListaDoblementeEnlazada struct {
	Cabeza, cola *Nodo
	Tamano       int
}

//Función que permite agregar las tiendas a la lista doblemente enlazada
func (Lista *ListaDoblementeEnlazada)AgregarAlPrincipio(nombre string, descripcion string, contacto string, calificacion int, Logo string) {
	n := &Nodo{Dato: Tienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion, Logo: Logo}}
	if Lista.Tamano == 0{
		Lista.Cabeza = n
		Lista.cola = n
	}else{
		aux := Lista.Cabeza
		Lista.Cabeza = n
		aux.Anterior = Lista.Cabeza
		Lista.Cabeza.Siguiente = aux
	}
	Lista.Tamano++
}

//Función que permite agregar las tiendas a la lista doblemente enlazada
func (Lista *ListaDoblementeEnlazada)AgregarAlFinal(nombre string, descripcion string, contacto string, calificacion int, Logo string){
	n := &Nodo{Dato: Tienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion, Logo: Logo}}
	if Lista.Tamano == 0 {
		Lista.Cabeza = n
		Lista.cola = n
	}else{
		aux := Lista.cola
		Lista.cola = n
		aux.Siguiente = Lista.cola
		Lista.cola.Anterior = aux
	}
	Lista.Tamano++
}

//Permite eliminar tiendas del sistema
func (Lista *ListaDoblementeEnlazada) Eliminar(valor string){
	if Lista.Cabeza.Dato.Nombre == valor{
		Lista.Cabeza = Lista.Cabeza.Siguiente
	} else if Lista.cola.Dato.Nombre == valor{
		Lista.cola = Lista.cola.Anterior
	}else{
		auxiliar := Lista.Cabeza
		for auxiliar.Siguiente.Dato.Nombre != valor{
			auxiliar = auxiliar.Siguiente
			if auxiliar.Siguiente == nil{
				fmt.Println("El valor a eliminar no existe")
				return
			}
		}
		auxiliar.Siguiente = auxiliar.Siguiente.Siguiente
	}
	Lista.Tamano--
}

//Funcion buscar
func (Lista *ListaDoblementeEnlazada) Search(valor string) (bool){
	if Lista.Cabeza.Dato.Nombre == valor{
		return true
	} else if Lista.cola.Dato.Nombre == valor{
		return true
	}else{
		auxiliar := Lista.Cabeza
		for auxiliar.Siguiente.Dato.Nombre != valor{
			auxiliar = auxiliar.Siguiente
			if auxiliar.Siguiente == nil{
				fmt.Println("El valor a eliminar no existe")
				return false
			}
		}
		return true
	}
}

//funcion buscar que retorna una tienda
func (Lista *ListaDoblementeEnlazada) Buscar(tienda TiendaIntroducida)(Tienda){
	if Lista.Cabeza.Dato.Nombre == tienda.Nombre{
		return Lista.Cabeza.Dato
	} else if Lista.cola.Dato.Nombre == tienda.Nombre{
		return Lista.cola.Dato
	}else{
		auxiliar := Lista.Cabeza
		for auxiliar.Siguiente.Dato.Nombre != tienda.Nombre{
			auxiliar = auxiliar.Siguiente
			if auxiliar.Siguiente == nil{
				return auxiliar.Siguiente.Dato
				fmt.Println("El valor a eliminar no existe")
			}
		}
		return auxiliar.Siguiente.Dato
	}
}

//Funcion para la imporesion de nodos
func (Lista ListaDoblementeEnlazada) Imprimir(){
	toPrint := Lista.Cabeza
	if Lista.Tamano == 0{
		fmt.Println("No hay Tienda")
		return
	}
	for Lista.Tamano != 0 {
		fmt.Println(toPrint.Dato.Nombre)
		toPrint = toPrint.Siguiente
		Lista.Tamano--
	}
}
