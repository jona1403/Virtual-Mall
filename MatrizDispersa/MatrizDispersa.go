package MatrizDispersa

type ProductoPedido struct{
	Codigo int
}

type Pedido struct{
	Fecha string
	Tienda string
	Departamento string
	Calificacion int
	Productos []ProductoPedido
}

type BD_Pedidos struct {
	Pedidos []Pedido
}

type NodoC struct{
	x, y int
	Arriba, Abajo, Izquierda, Derecha *NodoC
	Dato *Pedido
	header int
	siguiente, anterior *NodoC
}

type lista struct{
	cabeza, cola *NodoC
}

type matriz struct{
	cabeceraHor, vabeceraVer *lista

}

func (nodo *NodoC) NodoMatriz(x int, y int, pedido *Pedido) *NodoC{
	return &NodoC{x:x,y:y,Arriba:nil,Abajo:nil,Izquierda:nil,Derecha:nil,Dato:pedido,header:nil,siguiente: nil,anterior:  nil}
}

func NodoLista(header int) *NodoC{
	return &NodoC{x:nil,y:nil,Arriba:nil,Abajo:nil,Izquierda:nil,Derecha:nil,Dato:nil,header:header,siguiente:nil,anterior:nil}
}

func nuevaLista() *lista{
	return &lista{nil, nil}
}

func nuevaMatriz() *matriz{
	return &matriz{nuevaLista(), nuevaLista()}
}

func (n *NodoC) headerX() int{ return n.x }
func (n *NodoC) headerY() int{ return n.y }
func (n *NodoC) toString() string{return ""}

func (l *lista) Ordenar( nuevo *NodoC ){
	aux := l.cabeza
	for aux != nil {
		if nuevo.header > aux.header{
			aux= aux.siguiente
		}else{
			if aux == l.cabeza{
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.cabeza = nuevo
			}else{
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.cola.siguiente = nuevo
	nuevo.anterior = l.cola
	l.cola = nuevo
}

func (l *lista) Insertar(header int){
	 	nuevo := NodoLista(header)
	 	if l.cabeza == nil{
	 		l.cabeza = nuevo
	 		l.cola = nuevo
		}else{
			l.Ordenar(nuevo)
		}
}

func (l *lista) Buscar(header int) *NodoC{
	temporal := l.cabeza
	for temporal != nil{
		if temporal.header == header{
			return temporal
		}
		temporal = temporal.siguiente
	}
	return nil
}

func (m *matriz) insertMatr (x int, y int, pedido *Pedido){

}