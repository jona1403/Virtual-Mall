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
	Ordenador(db)
}

//Funcion de escritura del archivo JSON
func EncoderJson(){
	Tiendas := []ListaDoblementeEnlazada.Tienda{{Nombre: "Aurora", Descripcion: "Es una empresa multinacional " +
		"estadounidense dedicada al diseño, desarrollo, fabricación y comercialización " +
		"de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros " +
		"artículos deportivos", Contacto: "5544-3377", Calificacion: "5"}}

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

//Metodo de ordenamiento dentro de el vector de datos
func Ordenador(db DB_VirtualMall) {

	//var ArregloDepartamentos []map[int]Departamento

	//m["Deportes"] = Departamento{Nombre: ""}
	//var BD [][26][5] ListaDoblementeEnlazada
	for i:= 0; i<len(db.Datos); i++{
		for j:= 0; j<len(db.Datos[i].Departamentos); j++{
			for k:= 0; k<len(db.Datos[i].Departamentos[j].Tiendas); k++{
				fmt.Println("Hola")
			}
		}
	}

}
