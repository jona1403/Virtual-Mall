package MatrizDispersa

type node struct{
	matriz Matriz
	mes int
	Altura int
	izquierda, derecha *node
}

type Arbol struct{
	root *node
}

func (avl Arbol) altura(temporal *node) int{
	if temporal == nil{return -1}
	return temporal.Altura
}

func (avl Arbol) max(valor1 int, valor2 int) int{
	if valor1 > valor2{return valor1}
	return valor2
}

func (avl *Arbol) add(matriz Matriz, mes int){
	avl.root = avl.add2(matriz, mes, avl.root)
}

func (avl Arbol) add2(matriz Matriz, mes int, temporal *node) *node{
	if temporal == nil{
		return &node{matriz: matriz, mes: mes}
	}else if mes > temporal.mes{
		temporal.derecha = avl.add2(matriz, mes, temporal.derecha)
		if(avl.altura(temporal.derecha)-avl.altura(temporal.izquierda) == 2){
			if mes > temporal.derecha.mes{
				temporal = avl.srr(temporal)
			}else{
				temporal = avl.drr(temporal)
			}
		}
		//Verificar esta condicion
	}else{
		temporal.izquierda = avl.add2(matriz, mes,temporal.izquierda)
		if avl.altura(temporal.izquierda) - avl.altura(temporal.derecha) == 2{
			if mes < temporal.izquierda.mes{
				temporal = avl.srl(temporal)
			}else{
				temporal = avl.drl(temporal)
			}
		}
	}
	temporal.Altura = avl.max(avl.altura(temporal.derecha), avl.altura(temporal.izquierda))+1
	return temporal
}

func (avl Arbol) srl(temp *node) *node{
	temp2 := temp.izquierda
	temp.izquierda = temp2.derecha
	temp2.derecha = temp
	temp.Altura = avl.max(avl.altura(temp.derecha), avl.altura(temp.izquierda))+1
	temp2.Altura = avl.max(avl.altura(temp2.izquierda), temp.Altura)+1
	return temp2

}
func (avl Arbol) srr(temp *node) *node{
	temp2 := temp.derecha
	temp.derecha = temp2.izquierda
	temp2.izquierda = temp
	temp.Altura = avl.max(avl.altura(temp.derecha), avl.altura(temp.izquierda))+1
	temp2.Altura = avl.max(avl.altura(temp2.izquierda), temp.Altura)+1
	return temp2
}
func (avl Arbol) drl(temp *node) *node{
	temp.izquierda = avl.srr(temp.izquierda)
	return avl.srl(temp)
}
func (avl Arbol) drr(temp *node) *node{
	temp.derecha = avl.srl(temp.derecha)
	return avl.srr(temp)
}