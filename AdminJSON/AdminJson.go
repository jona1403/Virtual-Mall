package AdminJSON

import (
	"bytes"
	"encoding/json"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"

	"github.com/VirtualMall/ListaDoblementeEnlazada/ArbolAVL"
	"github.com/mzohreva/GoGraphviz/graphviz"
	"io"
	"log"
	"os"
	"strconv"
)

//Tipo datos indices, contiene el indice y los departamentos que se encuentran dentro de el
//corregido
type Datos_Indices struct {
	Indice string
	Departamentos []Departamentos
}
//corregido
type Departamentos struct{
	Nombre string
	Tiendas []ListaDoblementeEnlazada.Tienda
}

//Tipo db, dentro de este se encuentran todos los datos generales
//corregido
type DB_VirtualMall struct {
	Datos []Datos_Indices
}

//Funcion de escritura del archivo JSON
//corregido
func EncoderJson(db DB_VirtualMall){
	//Encoder
	var buf = new (bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(db)
	f, err := os.Create("Tiendas.json")
	if nil != err{
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}

//Metodo de obtencion de tamaño de departamentos y mapa departamentos
//corregido
func getMapaDepartamentos(db DB_VirtualMall)(map[int]string){
	MapaDepartamentos := make(map[int]string)
	var Aux int = 0

	for i:= 0; i<len(db.Datos); i++{
		for j:=0; j<len(db.Datos[i].Departamentos); j++{
			if !existeDepartamento(MapaDepartamentos,db.Datos[i].Departamentos[j].Nombre){
				MapaDepartamentos[Aux] = db.Datos[i].Departamentos[j].Nombre
				Aux++
			}
		}
	}
	return MapaDepartamentos
}

//Revisa la existencia del departamento
//corregido
func existeDepartamento(mapa map[int]string, valor string)(bool){
	for _, value := range mapa{

		if value == valor{
			return true
		}
	}
	return false
}

//Obtener la llave de los departamentos
func KeyDepto(mapa map[int]string, valor string)(int){
	for ke, valu := range mapa{
		if valu == valor{
			return ke
		}
	}
	return -1
}

//Obtener la llave de los departamentos
func KeyIndice(mapa map[int]string, valor string)(int){
	for ke, valu := range mapa{
		if valu == string(rune(valor[0])){
			return ke
		}
	}
	return -1
}


//Obtencion mapa indices
//Corregido
func getMapaIndices(db DB_VirtualMall)(map[int]string){
	MapaIndices := make(map[int]string)
	var Aux int = 0

	for i:= 0; i<len(db.Datos); i++{
		if !existeDepartamento(MapaIndices,db.Datos[i].Indice){
			MapaIndices[Aux] = db.Datos[i].Indice
			Aux++
		}

	}
	return MapaIndices
}

//Metodo de linealizacion de datos
//Corregido
func Linealizacion(db DB_VirtualMall) ([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, map[int]string, map[int]string) {
	MapaDepartamentos := getMapaDepartamentos(db)
	MapaIndices:= getMapaIndices(db)
	var Indice string
	var runes []rune
	MatrizTiendas := make([][][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaDepartamentos))
	for i := range MatrizTiendas {
		MatrizTiendas[i] = make([][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaIndices))
	}
	for i := range MatrizTiendas {
		for j:= range MatrizTiendas[i]{
			MatrizTiendas[i][j] = make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, 5)
		}
	}
	for i:=0; i< len(db.Datos); i++{
		for j:=0; j < len(db.Datos[i].Departamentos); j++{
			Indice = db.Datos[i].Indice
			runes = []rune(Indice)
			for k:=0; k < len(db.Datos[i].Departamentos[j].Tiendas); k++{
				if KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre) != -1{
					if (int(runes[0]) >= 65 && int(runes[0]) <= 78) || (int(runes[0]) >= 97 && int(runes[0]) <= 110){
						if int(runes[0]) >= 65 && int(runes[0]) <= 78{
							MatrizTiendas[KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre)][runes[0]-65][db.Datos[i].Departamentos[j].Tiendas[k].Calificacion-1].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion, db.Datos[i].Departamentos[j].Tiendas[k].Logo)
						}else{
							MatrizTiendas[KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre)][runes[0]-97][db.Datos[i].Departamentos[j].Tiendas[k].Calificacion-1].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion, db.Datos[i].Departamentos[j].Tiendas[k].Logo)
						}
					}else if Indice == "ñ" || Indice == "Ñ"{
						MatrizTiendas[KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre)][14][db.Datos[i].Departamentos[j].Tiendas[k].Calificacion-1].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion, db.Datos[i].Departamentos[j].Tiendas[k].Logo)
					}else if (int(runes[0]) >= 79 && int(runes[0]) <= 90) || (int(runes[0]) >= 111 && int(runes[0]) <= 122){
						if int(runes[0]) >= 79 && int(runes[0]) <= 90{
							MatrizTiendas[KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre)][runes[0]-64][db.Datos[i].Departamentos[j].Tiendas[k].Calificacion-1].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion, db.Datos[i].Departamentos[j].Tiendas[k].Logo)
						}else{
							MatrizTiendas[KeyDepto(MapaDepartamentos, db.Datos[i].Departamentos[j].Nombre)][runes[0]-96][db.Datos[i].Departamentos[j].Tiendas[k].Calificacion-1].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion, db.Datos[i].Departamentos[j].Tiendas[k].Logo)
						}
					}
				}
			}
		}
	}

	ListaLinealizada := make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaDepartamentos)*len(MapaIndices)*5)



	for i:=0; i < len(MapaDepartamentos); i++{
		for j:=0; j < len(MapaIndices); j++{
			for k:=0; k < 5; k++{
				ListaLinealizada[i+len(MapaDepartamentos)*(j+len(MapaIndices)*k)] = MatrizTiendas[i][j][k]
			}
		}
	}
	return ListaLinealizada, MapaDepartamentos, MapaIndices
}

//Metodo de llenado de arboles
func AgregarProducto(db ArbolAVL.BD_Inventarios,lista []ListaDoblementeEnlazada.ListaDoblementeEnlazada,mapadepartamentos map[int]string,mapaindices map[int]string)([]ListaDoblementeEnlazada.ListaDoblementeEnlazada){

	for i:=0;i < len(db.Invetarios); i++{
		keydp := KeyDepto(mapadepartamentos, db.Invetarios[i].Departamento)
		keyin := KeyIndice(mapaindices, db.Invetarios[i].Tienda)
		for j:= 0; j< len(db.Invetarios[i].Productos); j++{
			lista[keydp+len(mapadepartamentos)*(keyin+len(mapaindices)*(db.Invetarios[i].Calificacion-1))].AgregarProducto(db.Invetarios[i].Productos[j], db.Invetarios[i].Tienda)
		}
		//s:= lista[keydp+len(mapadepartamentos)*(keyin+len(mapaindices)*(db.Invetarios[i].Calificacion-1))].GetArbol(db.Invetarios[i].Tienda)

	}
	return lista
}

//Grafica de lista linealizada
func Graficar(l []ListaDoblementeEnlazada.ListaDoblementeEnlazada){
	G := &graphviz.Graph{}
	nodos1:= make([]int, 0)
	nodos2:= make([]int, 0)
	for i:=0; i< len(l); i++{
		nodos1 = append(nodos1, G.AddNode(strconv.Itoa (i)))
		if l[i].Cabeza!= nil{
			cabeza:= l[i].Cabeza
			for cabeza!= nil {
				nodos2 = append(nodos2, G.AddNode(cabeza.Dato.Nombre))
				cabeza = cabeza.Siguiente
			}
		}
	}
	var aux1 int = 0
	var aux2 int = 0
	for i := 0; i < len(l); i++ {
		if l[i].Cabeza != nil{
			if aux1 == 0{
				G.AddEdge(nodos1[i], nodos2[aux2], "")
				G.AddEdge(nodos2[aux2], nodos1[i], "")
			}
			for j := 0; j < l[i].Tamano-1; j++ {
				G.AddEdge(nodos2[aux2], nodos2[aux2+1], "")
				G.AddEdge(nodos2[aux2+1], nodos2[aux2], "")
				aux2++
			}
			aux1++
			aux2++
		}
		aux1 = 0
	}
	for i := 0; i < len(nodos1)-1; i++ {
		G.AddEdge(nodos1[i], nodos1[i+1], "")
	}
	G.MakeSameRank(nodos2[0], nodos2[1], nodos2[2:]...)
	G.MakeSameRank(nodos1[0], nodos1[1], nodos1[2:]...)

	G.DefaultNodeAttribute(graphviz.Shape, graphviz.ShapeBox)
	G.DefaultNodeAttribute(graphviz.FontName, "Courier")
	G.DefaultEdgeAttribute(graphviz.FontName, "Courier")
	G.GraphAttribute(graphviz.NodeSep, "0.2")
	G.SetTitle("\n\n" + "Arreglo linealizado")
	G.MakeDirected()

	err := G.GenerateImage("dot", "lista.png", "png")
	if err != nil {
		log.Fatal(err)
	}
}

