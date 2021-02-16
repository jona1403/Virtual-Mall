package AdminJSON

import (
	"bytes"
	"encoding/json"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"
	"io"
	"log"
	"os"
)

//Tipo departamento, contiene el nombre del departamento y las tiendas que se encuentran dentro
//del el
type Departamento struct {
	Nombre string
	Tiendas []ListaDoblementeEnlazada.Tienda
}


//Tipo datos indices, contiene el indice y los departamentos que se encuentran dentro de el
type Datos_Indices struct {
	Indice string
	Departamentos []Departamento
}

//Tipo db, dentro de este se encuentran todos los datos generales
type DB_VirtualMall struct {
	Datos []Datos_Indices
}

//Metodo de lectura del archivo JSON
func DecoderJSON(RutaArchivo string){
	f, err := os.Open(RutaArchivo)
	if nil != err{
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	db:= DB_VirtualMall{}
	dec.Decode(&db)
}

//Funcion de escritura del archivo JSON
func EncoderJson(){
	Tiendas := []ListaDoblementeEnlazada.Tienda{{Nombre: "Aurora", Descripcion: "Es una empresa multinacional " +
		"estadounidense dedicada al diseño, desarrollo, fabricación y comercialización " +
		"de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros " +
		"artículos deportivos", Contacto: "5544-3377", Calificacion: 5}}

	Departamentos := []Departamento{{Nombre: "Deportes", Tiendas: Tiendas}}
	Datos2 := []Datos_Indices{{Indice: "A", Departamentos: Departamentos}}
	DB := DB_VirtualMall{Datos: Datos2}
	//Encoder
	var buf = new (bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(DB)
	f, err := os.Create("user.db.json")
	if nil != err{
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}

//Metodo de obtencion de cantidad de departamentos en la bd
func getCantidadDepartamentos(db DB_VirtualMall) (map[string]int){
	MapaDepartamentos := make(map[string]int)
	var CantidadDeptos int = 0
	for i:= 0; i<len(db.Datos); i++{
		for j:= CantidadDeptos; j < len(db.Datos[i].Departamentos)+CantidadDeptos; j++{
			MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre] = j
		}
		CantidadDeptos += len(db.Datos[i].Departamentos)
	}

	return MapaDepartamentos
}

//Función de ordenamiento de datos mediante col row
func OrdenarTiendas(db DB_VirtualMall){
	var Indice string
	var runes []rune
	MapaDepartamentos := getCantidadDepartamentos(db)
	//const cantidad = len(MapaDepartamentos)
	var MatrizTiendas [30][27][5]ListaDoblementeEnlazada.ListaDoblementeEnlazada
	for i:= 0; i<len(db.Datos); i++{
		for j:= 0; j< len(db.Datos[i].Departamentos); j++{
			Indice = db.Datos[i].Indice
			runes = []rune(Indice)
			for k:= 0; k < len(db.Datos[i].Departamentos[j].Tiendas); k++{
				if (int(runes[0]) >= 65 && int(runes[0]) <= 78) || (int(runes[0]) >= 97 && int(runes[0]) <= 110){
					if int(runes[0]) >= 65 && int(runes[0]) <= 78{
						MatrizTiendas[MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre]][runes[0]-65][k].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
					}else{
						MatrizTiendas[MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre]][runes[0]-97][k].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
					}
				}else if Indice == "ñ" || Indice == "Ñ"{
					MatrizTiendas[MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre]][14][k].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
				}else if (int(runes[0]) >= 79 && int(runes[0]) <= 90) || (int(runes[0]) >= 111 && int(runes[0]) <= 122){
					if int(runes[0]) >= 79 && int(runes[0]) <= 90{
						MatrizTiendas[MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre]][runes[0]-64][k].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
					}else{
						MatrizTiendas[MapaDepartamentos[db.Datos[i].Departamentos[j].Nombre]][runes[0]-96][k].AgregarAlPrincipio(db.Datos[i].Departamentos[j].Tiendas[k].Nombre, db.Datos[i].Departamentos[j].Tiendas[k].Descripcion, db.Datos[i].Departamentos[j].Tiendas[k].Contacto, db.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
					}
				}
			}
		}
	}
}

func LinealizarMatriz(){

}