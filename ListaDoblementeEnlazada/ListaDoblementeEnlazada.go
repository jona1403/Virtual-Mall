package ListaDoblementeEnlazada

import "fmt"

//Tipo tienda, tiene todos los atributos de las tiendas
type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
}

//Nodo, es utilizado para la lista doblemente enlazada
type Nodo struct {
	anterior,siguiente *Nodo
	Dato Tienda
}

//Lista doblemente enlazada, contiene todas las tiendas que se encuentren dentro
//de la misma sección y mismo puntaje
type ListaDoblementeEnlazada struct {
	cabeza, cola *Nodo
	tamano int
}

//Función que permite agregar las tiendas a la lista doblemente enlazada
func (Lista *ListaDoblementeEnlazada)AgregarAlPrincipio(nombre string, descripcion string, contacto string, calificacion int) {
	n := &Nodo{Dato: Tienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion}}
	if Lista.tamano == 0{
		Lista.cabeza = n
		Lista.cola = n
	}else{
		aux := Lista.cabeza
		Lista.cabeza = n
		aux.anterior = Lista.cabeza
		Lista.cabeza.siguiente = aux
	}
	Lista.tamano ++
}

//Función que permite agregar las tiendas a la lista doblemente enlazada
func (Lista *ListaDoblementeEnlazada)AgregarAlFinal(nombre string, descripcion string, contacto string, calificacion int){
	n := &Nodo{Dato: Tienda{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion}}
	if Lista.tamano == 0 {
		Lista.cabeza = n
		Lista.cola = n
	}else{
		aux := Lista.cola
		Lista.cola = n
		aux.siguiente = Lista.cola
		Lista.cola.anterior = aux
	}
	Lista.tamano ++
}

//Permite eliminar tiendas del sistema
func (Lista *ListaDoblementeEnlazada) Eliminar(valor string){
	if Lista.cabeza.Dato.Nombre == valor{
		Lista.cabeza = Lista.cabeza.siguiente
	} else if Lista.cola.Dato.Nombre == valor{
		Lista.cola = Lista.cola.anterior
	}else{
		auxiliar := Lista.cabeza
		for auxiliar.siguiente.Dato.Nombre != valor{
			auxiliar = auxiliar.siguiente
			if auxiliar.siguiente == nil{
				fmt.Println("El valor a eliminar no existe")
				return
			}
		}
		auxiliar.siguiente = auxiliar.siguiente.siguiente
	}
	Lista.tamano --
}

//Funcion para la imporesion de nodos
func (Lista ListaDoblementeEnlazada) Imprimir(){
	toPrint := Lista.cabeza
	if Lista.tamano == 0{
		fmt.Println("No hay Tienda")
		return
	}
	for Lista.tamano != 0 {
		fmt.Println(toPrint.Dato.Nombre)
		toPrint = toPrint.siguiente
		Lista.tamano --
	}
}
