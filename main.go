package main

import (
	"fmt"
)

func main(){
	List := ListaDoblementeEnlazada{}
	node1:= &Nodo{Dato: Tienda{nombre: "nodo1"}}
	node2:= &Nodo{Dato: Tienda{nombre: "nodo2"}}
	node3:= &Nodo{Dato: Tienda{nombre: "nodo3"}}
	node4:= &Nodo{Dato: Tienda{nombre: "nodo4"}}
	List.AgregarAlFinal(node1)
	List.AgregarAlFinal(node2)
	List.AgregarAlFinal(node3)
	List.AgregarAlFinal(node4)
	List.Eliminar("nodo6")
	List.Imprimir()
}

//Tipo tienda, tiene todos los atributos de las tiendas
type Tienda struct {
	nombre string
	descripcion string
	contacto string
	calificacion int
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
func (Lista *ListaDoblementeEnlazada)AgregarAlPrincipio(n *Nodo) {
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
func (Lista *ListaDoblementeEnlazada)AgregarAlFinal(n *Nodo){
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
	if Lista.cabeza.Dato.nombre == valor{
		Lista.cabeza = Lista.cabeza.siguiente
	} else if Lista.cola.Dato.nombre == valor{
		Lista.cola = Lista.cola.anterior
	}else{
		auxiliar := Lista.cabeza
		for auxiliar.siguiente.Dato.nombre != valor{
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

func (Lista ListaDoblementeEnlazada) Imprimir(){
	toPrint := Lista.cabeza
	if Lista.tamano == 0{
		fmt.Println("No hay Tienda")
		return
	}
	for Lista.tamano != 0 {
		fmt.Println(toPrint.Dato.nombre)
		toPrint = toPrint.siguiente
		Lista.tamano --
	}
}
