package MatrizDispersa

import "strconv"

func agregarpedidosmatriz(db BD_Pedidos, arbanio Arbolanio, mapadepatamentos map[int]string) (Arbolanio){
	for i:= 0; i < len(db.Pedidos); i++{
		s:= db.Pedidos[i].Fecha
		//dia, _ := strconv.Atoi(s[0:2])
		//mes, _ := strconv.Atoi(s[3:5])
		anio, _ := strconv.Atoi(s[6:])
		if(arbanio.SearchAnios(anio, arbanio.root)){
			arbanio.AddmonthsToAnios(anio, db.Pedidos[i])
		}else{
			arbanio.add(Arbol{}, anio)
		}
	}
	return arbanio
}
