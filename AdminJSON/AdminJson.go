package AdminJSON

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"
	"io"
	"log"
	"os"
)

//Tipo datos indices, contiene el indice y los departamentos que se encuentran dentro de el
type Datos_Indices struct {
	Indice string
	Departamentos []map[string][]ListaDoblementeEnlazada.Tienda
}

//Tipo db, dentro de este se encuentran todos los datos generales
type DB_VirtualMall struct {
	Datos []Datos_Indices
}

//Funcion de escritura del archivo JSON
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
func getMapaDepartamentos(db DB_VirtualMall)(map[int]string){
	MapaDepartamentos := make(map[int]string)
	var Aux int = 0
	for i:=0; i<len(db.Datos);i++{
		for j:=0; j<len(db.Datos[i].Departamentos);j++{
			for key, _ := range db.Datos[i].Departamentos[j]{
				if !existeDepartamento(MapaDepartamentos,key){
					MapaDepartamentos[Aux] = key
					Aux++
				}
			}
		}
	}
	return MapaDepartamentos
}

//Revisa la existencia del departamento
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

//Metodo de linealizacion de datos
func Linealizacion(db DB_VirtualMall) ([]ListaDoblementeEnlazada.ListaDoblementeEnlazada,map[int]string) {
	MapaDepartamentos := getMapaDepartamentos(db)
	var Indice string
	var runes []rune
	MatrizTiendas := make([][][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaDepartamentos))
	for i := range MatrizTiendas {
		MatrizTiendas[i] = make([][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, 27)
	}
	for i := range MatrizTiendas {
		for j:= range MatrizTiendas[i]{
			MatrizTiendas[i][j] = make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, 5)
		}
	}
	for i:=0; i<len(db.Datos);i++{
		for j:=0; j<len(db.Datos[i].Departamentos);j++{
			for key, value := range db.Datos[i].Departamentos[j]{
				Indice = db.Datos[i].Indice
				runes = []rune(Indice)
				for _, valu := range value{
					if KeyDepto(MapaDepartamentos, key) != -1{
						if (int(runes[0]) >= 65 && int(runes[0]) <= 78) || (int(runes[0]) >= 97 && int(runes[0]) <= 110){
							if int(runes[0]) >= 65 && int(runes[0]) <= 78{
								MatrizTiendas[KeyDepto(MapaDepartamentos, key)][runes[0]-65][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
							}else{
								MatrizTiendas[KeyDepto(MapaDepartamentos, key)][runes[0]-97][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
							}
						}else if Indice == "ñ" || Indice == "Ñ"{
							MatrizTiendas[KeyDepto(MapaDepartamentos, key)][14][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
						}else if (int(runes[0]) >= 79 && int(runes[0]) <= 90) || (int(runes[0]) >= 111 && int(runes[0]) <= 122){
							if int(runes[0]) >= 79 && int(runes[0]) <= 90{
								MatrizTiendas[KeyDepto(MapaDepartamentos, key)][runes[0]-64][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
							}else{
								MatrizTiendas[KeyDepto(MapaDepartamentos, key)][runes[0]-96][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
							}
						}
					}

				}
			}
		}
	}
	ListaLinealizada := make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaDepartamentos)*27*5)
	fmt.Println(len(ListaLinealizada))
	for i:=0; i < len(MapaDepartamentos); i++{
		for j:=0; j < 27; j++{
			for k:=0; k < 5; k++{
				ListaLinealizada[i+len(MapaDepartamentos)*(j+27*k)] = MatrizTiendas[i][j][k]
			}
		}
	}
	return ListaLinealizada, MapaDepartamentos
}

//Grafica de lista linealizada
func Graficar([]ListaDoblementeEnlazada.ListaDoblementeEnlazada){}