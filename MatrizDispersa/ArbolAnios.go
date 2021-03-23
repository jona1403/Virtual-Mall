package MatrizDispersa

type nodeanio struct{
	arbolmeses             Arbol
	anio               int
	Altura             int
	izquierda, derecha *nodeanio
}

type Arbolanio struct{
	root *nodeanio
}

func (avl Arbolanio) altura(temporal *nodeanio) int{
	if temporal == nil{return -1}
	return temporal.Altura
}

func (avl Arbolanio) max(valor1 int, valor2 int) int{
	if valor1 > valor2{return valor1}
	return valor2
}

func (avl *Arbolanio) add(arbol Arbol, mes int){
	avl.root = avl.add2(arbol, mes, avl.root)
}

func (avl Arbolanio) add2(arbol Arbol, anio int, temporal *nodeanio) *nodeanio {
	if temporal == nil{
		return &nodeanio{arbolmeses: arbol, anio: anio}
	}else if anio > temporal.anio {
		temporal.derecha = avl.add2(arbol, anio, temporal.derecha)
		if(avl.altura(temporal.derecha)-avl.altura(temporal.izquierda) == 2){
			if anio > temporal.derecha.anio {
				temporal = avl.srr(temporal)
			}else{
				temporal = avl.drr(temporal)
			}
		}
		//Verificar esta condicion
	}else{
		temporal.izquierda = avl.add2(arbol, anio,temporal.izquierda)
		if avl.altura(temporal.izquierda) - avl.altura(temporal.derecha) == 2{
			if anio < temporal.izquierda.anio {
				temporal = avl.srl(temporal)
			}else{
				temporal = avl.drl(temporal)
			}
		}
	}
	temporal.Altura = avl.max(avl.altura(temporal.derecha), avl.altura(temporal.izquierda))+1
	return temporal
}

func (avl Arbolanio) srl(temp *nodeanio) *nodeanio {
	temp2 := temp.izquierda
	temp.izquierda = temp2.derecha
	temp2.derecha = temp
	temp.Altura = avl.max(avl.altura(temp.derecha), avl.altura(temp.izquierda))+1
	temp2.Altura = avl.max(avl.altura(temp2.izquierda), temp.Altura)+1
	return temp2

}
func (avl Arbolanio) srr(temp *nodeanio) *nodeanio {
	temp2 := temp.derecha
	temp.derecha = temp2.izquierda
	temp2.izquierda = temp
	temp.Altura = avl.max(avl.altura(temp.derecha), avl.altura(temp.izquierda))+1
	temp2.Altura = avl.max(avl.altura(temp2.izquierda), temp.Altura)+1
	return temp2
}
func (avl Arbolanio) drl(temp *nodeanio) *nodeanio {
	temp.izquierda = avl.srr(temp.izquierda)
	return avl.srl(temp)
}
func (avl Arbolanio) drr(temp *nodeanio) *nodeanio {
	temp.derecha = avl.srl(temp.derecha)
	return avl.srr(temp)
}
