package AdminJSON

import (
	"encoding/json"
	"fmt"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"
	"log"
	"os"
)

//Tipo departamento, contiene el nombre del departamento y las tiendas que se encuentran dentro
//del el

//Tipo datos indices, contiene el indice y los departamentos que se encuentran dentro de el
type Datos_Indices struct {
	Indice string
	Departamentos []map[string][]ListaDoblementeEnlazada.Tienda
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
	Linealizacion(db)
}

func UnmarshallJSON(JSON string){

	Json := `{
  "Datos": [{
    "Indice": "A",
    "Departamentos": [{
      "Deportes": [{
        "Nombre": "Aurora",
        "Descripcion": "Es una empresa multinacional estadounidense dedicada al diseño, desarrollo, fabricación y comercialización de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros artículos deportivos",
        "Contacto": "5544-3377",
        "Calificacion": 5
      },
        {
          "Nombre": "Amador",
          "Descripcion": "es una empresa alemana fabricante de accesorios, ropa y calzado deportivo, cuya sede central está en Herzogenaurach, Alemania",
          "Contacto": "5588-9988",
          "Calificacion": 4
        },
        {
          "Nombre": "Armados",
          "Descripcion": "Equipo extremo",
          "Contacto": "8995222",
          "Calificacion": 5
        }
      ]
    },
      {
        "Comida": [{
          "Nombre": "A comer todo",
          "Descripcion": "todo lo que puedan pedir por un dollar",
          "Contacto": "559999",
          "Calificacion": 5
        }]
      },
      {
        "Celulares": []
      }
    ]
  },
    {
      "Indice": "B",
      "Departamentos": [{
        "Deportes": []
      },
        {
          "Comida": []
        },
        {
          "Celulares": [{
            "Nombre": "Bayoneta",
            "Descripcion": "Telefonos militares",
            "Contacto": "bayoneta@gmail.com",
            "Calificacion": 5
          }]
        }
      ]
    }
  ]
}`
	var db DB_VirtualMall
	json.Unmarshal([]byte(Json), &db)

	for i:=0; i<len(db.Datos);i++{
		for j:=0; j<len(db.Datos[i].Departamentos);j++{
			for key, value := range db.Datos[i].Departamentos[j]{
				fmt.Println(key, value)
				for ke, valu := range value{
					fmt.Println(ke, valu)

				}
			}
		}
	}
}

//Funcion de escritura del archivo JSON
func EncoderJson(){
	/*Tiendas := []ListaDoblementeEnlazada.Tienda{{Nombre: "Aurora", Descripcion: "Es una empresa multinacional " +
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
	io.Copy(f, buf)*/
}

//Metodo de obtencion de tamaño de departamentos y mapa departamentos
func getMapaDepartamentos(db DB_VirtualMall)(map[int]string){
	MapaDepartamentos := make(map[int]string)
	var Aux int = 0
	for i:=0; i<len(db.Datos);i++{
		for j:=0; j<len(db.Datos[i].Departamentos);j++{
			for key, value := range db.Datos[i].Departamentos[j]{
				fmt.Println(key, value)
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
	for key, value := range mapa{
		fmt.Println(key, value)
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
func Linealizacion(db DB_VirtualMall){
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
				fmt.Println(key, value)
				for ke, valu := range value{
					fmt.Println(ke)
					if (int(runes[0]) >= 65 && int(runes[0]) <= 78) || (int(runes[0]) >= 97 && int(runes[0]) <= 110){
						if int(runes[0]) >= 65 && int(runes[0]) <= 78{
							if KeyDepto(MapaDepartamentos, key) == -1{

							}else{
								MatrizTiendas[KeyDepto(MapaDepartamentos, key)][runes[0]-65][valu.Calificacion-1].AgregarAlPrincipio(valu.Nombre, valu.Descripcion, valu.Contacto, valu.Calificacion)
							}

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
	MatrizTiendas[2][1][4].Imprimir()
	ListaLinealizada := make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(MapaDepartamentos)*27*5)
	for i:=0; i < len(MapaDepartamentos); i++{
		for j:=0; j < 27; j++{
			for k:=0; k < 5; k++{
				ListaLinealizada[k+2*(j+2*i)] = MatrizTiendas[i][j][k]
			}
		}
	}
}