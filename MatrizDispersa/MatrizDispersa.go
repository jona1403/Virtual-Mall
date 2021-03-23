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

type Matriz struct{
	cabeceraHor, cabeceraVer *lista

}

func NodoMatriz(x int, y int, pedido *Pedido) *NodoC{
	return &NodoC{x:x,y:y,Dato:pedido}
}

func NodoLista(header int) *NodoC{
	return &NodoC{header:header}
}

func nuevaLista() *lista{
	return &lista{nil, nil}
}

func nuevaMatriz() *Matriz {
	return &Matriz{nuevaLista(), nuevaLista()}
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

func (m *Matriz) insertMatr (x int, y int, pedido *Pedido){
	horizontal := m.cabeceraHor.Buscar(x)
	vertical := m.cabeceraVer.Buscar(y)

	if horizontal == nil && vertical == nil{
		m.Noexisteninguno(x, y, pedido)
	}else if horizontal == nil && vertical != nil{
		m.SoloVertical(x, y, pedido)
	}else if vertical == nil && horizontal != nil{
		m.SoloHorizontal(x, y, pedido)
	}else{
		m.ExistenAmbas(x, y, pedido)
	}
}

func (m *Matriz) Noexisteninguno(x int, y int, pedido *Pedido){
	m.cabeceraHor.Insertar(x)
	m.cabeceraVer.Insertar(y)

	horizontal := m.cabeceraHor.Buscar(x)
	vertical := m.cabeceraVer.Buscar(y)

	nuevo := NodoMatriz(x, y , pedido)

	horizontal.Abajo = nuevo
	nuevo.Arriba = horizontal

	vertical.Derecha = nuevo
	nuevo.Izquierda = vertical

}

func (m *Matriz) SoloVertical(x int, y int, pedido *Pedido){
	m.cabeceraHor.Insertar(x)
	horizontal := m.cabeceraHor.Buscar(x)
	vertical := m.cabeceraVer.Buscar(y)

	nuevo := NodoMatriz(x, y, pedido)
	isagregado := false

	aux:= vertical.Derecha

	var cabecera int

	for aux != nil{
		cabecera = aux.headerX()
		if cabecera < x{
			aux = aux.Derecha
		}else{
			nuevo.Derecha = aux
			nuevo.Izquierda = aux.Izquierda
			aux.Izquierda.Derecha = nuevo
			aux.Izquierda = nuevo
			isagregado = true
			break
		}
	}
	if !isagregado {
		aux = vertical.Derecha
		for aux.Derecha != nil{
			aux = aux.Derecha
		}
		nuevo.Izquierda = aux
		aux.Derecha = nuevo
	}

	nuevo.Arriba = horizontal
	horizontal.Abajo = nuevo
}

func (m *Matriz) SoloHorizontal(x int, y int, pedido *Pedido){
	m.cabeceraVer.Insertar(y)
	horizontal := m.cabeceraHor.Buscar(x)
	vertical := m.cabeceraVer.Buscar(y)

	nuevo := NodoMatriz(x, y, pedido)
	isagregado := false

	aux:= horizontal.Abajo

	var cabecera int

	for aux != nil{
		cabecera = aux.headerY()
		if cabecera < y{
			aux = aux.Abajo
		}else{
			nuevo.Abajo = aux
			nuevo.Arriba = aux.Arriba
			aux.Arriba.Abajo = nuevo
			aux.Arriba = nuevo
			isagregado = true
			break
		}
	}
	if !isagregado {
		aux = horizontal.Abajo
		for aux.Abajo != nil{
			aux = aux.Abajo
		}
		nuevo.Arriba = aux
		aux.Abajo = nuevo
	}

	nuevo.Izquierda = vertical
	vertical.Derecha = nuevo
}

func (m *Matriz) ExistenAmbas(x int, y int, pedido *Pedido){
	horizontal := m.cabeceraHor.Buscar(x)
	vertical := m.cabeceraVer.Buscar(y)

	nuevo := NodoMatriz(x, y, pedido)
	isagregado := false

	aux:= vertical.Derecha

	var cabecera int

	for aux != nil{
		cabecera = aux.headerX()
		if cabecera < x{
			aux = aux.Derecha
		}else{
			nuevo.Derecha = aux
			nuevo.Izquierda = aux.Izquierda
			aux.Izquierda.Derecha = nuevo
			aux.Izquierda = nuevo
			isagregado = true
			break
		}
	}
	if !isagregado {
		aux = vertical.Derecha
		for aux.Derecha != nil{
			aux = aux.Derecha
		}
		nuevo.Izquierda = aux
		aux.Derecha = nuevo
	}

	isagregado = false

	aux = horizontal.Abajo



	for aux != nil{
		cabecera = aux.headerY()
		if cabecera < y{
			aux = aux.Abajo
		}else{
			nuevo.Abajo = aux
			nuevo.Arriba = aux.Arriba
			aux.Arriba.Abajo = nuevo
			aux.Arriba = nuevo
			isagregado = true
			break
		}
	}
	if !isagregado {
		aux = horizontal.Abajo
		for aux.Abajo != nil{
			aux = aux.Abajo
		}
		nuevo.Arriba = aux
		aux.Abajo = nuevo
	}
}
