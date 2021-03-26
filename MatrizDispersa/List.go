package MatrizDispersa

type ListNode struct{
	Siguiente, Anterior *ListNode
	Dato Pedido
}

type List struct{
	cabeza, cola *ListNode
	tamanio int
}

func(C *List)add(ped Pedido){
	n := &ListNode{Dato: ped}
	if C.tamanio == 0{
		C.cabeza.Dato = ped
		C.cola.Dato = ped
	}else{
		aux := C.cola
		C.cola = n
		aux.Siguiente = C.cola
		C.cola.Anterior = aux
	}
	C.tamanio++
}