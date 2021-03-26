package ArbolAVL

import "fmt"

type Producto struct {
	Nombre string
	Codigo int
	Descripcion string
	Precio int
	Cantidad int
	Imagen string
}

type InventarioTienda struct{
	Tienda string
	Departamento string
	Calificacion int
	Productos []Producto
}

type BD_Inventarios struct {
	Invetarios []InventarioTienda
}

//Tipo nodo del arbol AVL
type AVLnode struct {
	Product            Producto
	Altura             int
	Izquierdo, Derecho *AVLnode
}

//Tipo Arbol AVL
type AVLTree struct{
	root *AVLnode
}

//Funcion de retorno de altura del arbol
func (avl AVLTree) altura(temporal *AVLnode) int{
	if temporal == nil{return -1}
	return temporal.Altura
}

func (avl AVLTree) max(valor1 int, valor2 int) int{
	if valor1 > valor2{return valor1}
	return valor2
}

//Funcion Agregar
func (avl *AVLTree) Agegar(product Producto){
	avl.root = avl.Agregar2(product, avl.root)
}

//Funcion agregar para la sobre carga de datos
func (avl AVLTree) Agregar2(producto Producto, temporal *AVLnode) *AVLnode{
	if temporal == nil{
		return &AVLnode{Product: producto}
	}else if producto.Codigo > temporal.Product.Codigo{
		temporal.Derecho = avl.Agregar2(producto, temporal.Derecho)
		if(avl.altura(temporal.Derecho)-avl.altura(temporal.Izquierdo) == 2){
			if producto.Codigo > temporal.Derecho.Product.Codigo{
				temporal = avl.srr(temporal)
			}else{
				temporal = avl.drr(temporal)
			}
		}
	//Verificar esta condicion
	}else if producto.Codigo == temporal.Product.Codigo{
		temporal.Product.Cantidad+= producto.Cantidad
		return temporal
	}else{
		temporal.Izquierdo = avl.Agregar2(producto, temporal.Izquierdo)
		if avl.altura(temporal.Izquierdo) - avl.altura(temporal.Derecho) == 2{
			if producto.Codigo < temporal.Izquierdo.Product.Codigo{
				temporal = avl.srl(temporal)
			}else{
				temporal = avl.drl(temporal)
			}
		}
	}
	temporal.Altura = avl.max(avl.altura(temporal.Derecho), avl.altura(temporal.Izquierdo))+1
	return temporal
}

func (avl *AVLTree) RestarUnidades(codigo int, cantidad int){

}

//Rotacion simple a la derecha
func (avl AVLTree) srl(temporal *AVLnode) *AVLnode{
	temporal2 := temporal.Izquierdo
	temporal.Izquierdo = temporal2.Derecho
	temporal2.Derecho = temporal
	temporal.Altura = avl.max(avl.altura(temporal.Derecho), avl.altura(temporal.Izquierdo))+1
	temporal2.Altura = avl.max(avl.altura(temporal2.Izquierdo), temporal.Altura)+1
	return temporal2
}

//Rotacion simple por la izquierda
func (avl AVLTree) srr(temporal *AVLnode) *AVLnode{
	temporal2:= temporal.Derecho
	temporal.Derecho = temporal2.Izquierdo
	temporal2.Izquierdo = temporal
	temporal.Altura = avl.max(avl.altura(temporal.Derecho), avl.altura(temporal.Izquierdo))+1
	temporal2.Altura = avl.max(avl.altura(temporal2.Izquierdo), temporal.Altura)+1
	return temporal2
}

//Rotacion doble por la derecha
func (avl AVLTree) drl(temporal *AVLnode) *AVLnode{
	temporal.Izquierdo = avl.srr(temporal.Izquierdo)
	return avl.srl(temporal)
}

//Rotacion doble por la izquierda
func (avl AVLTree) drr(temporal *AVLnode) *AVLnode{
	temporal.Derecho = avl.srl(temporal.Derecho)
	return avl.srr(temporal)
}

//Funcion de impresion de nodos en orden
func (avl AVLTree) enorden(temporal *AVLnode){
	if temporal != nil{
		avl.enorden(temporal.Izquierdo)
		fmt.Print(temporal.Product.Codigo)
		avl.enorden(temporal.Derecho)
	}
}

//Funcion buscar(Falta probarla)
func (avl AVLTree) Buscar(producto Producto , temporal *AVLnode) bool{
	if temporal != nil{
		if producto.Codigo == temporal.Product.Codigo{
			return true
		}else if producto.Codigo > temporal.Product.Codigo{
			return avl.Buscar(producto, temporal.Derecho)
		}else{
			return avl.Buscar(producto, temporal.Izquierdo)
		}
	}
	return false
}