package main

import (
	"encoding/json"
	"fmt"
	"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"net/http"
	//"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
)
var db AdminJSON.DB_VirtualMall
var lista []ListaDoblementeEnlazada.ListaDoblementeEnlazada
var departamentos map[int]string

//Muestra de datos
func getJSON(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}

//Carga de datos
func setJSON(w http.ResponseWriter, r *http.Request){
	Tiendas, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Datos no validos")
	}
	json.Unmarshal([]byte(Tiendas), &db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	lista, departamentos = AdminJSON.Linealizacion(db)

	json.NewEncoder(w).Encode("Datos cargados")
}

//Obtencion de arreglo grafico
func getArreglo(w http.ResponseWriter, r *http.Request){
	AdminJSON.Graficar(lista)
	json.NewEncoder(w).Encode("Grafica generada")
}

//Elimia un elemento
func eliminar(w http.ResponseWriter, r *http.Request){
	var aux int
	var tienda ListaDoblementeEnlazada.TiendaEliminar
	Tienda, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Datos no validos")
	}
	json.Unmarshal([]byte(Tienda), &tienda)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	key := AdminJSON.KeyDepto(departamentos, tienda.Categoria)
	if key != -1{
		aux = getKeyLetra(tienda.Nombre)
		if aux != -1{
			del(tienda)
			lista[key+len(departamentos)*(aux+27*(tienda.Calificacion-1))].Eliminar(tienda.Nombre)
			json.NewEncoder(w).Encode("OK")
		}
	}else{
		json.NewEncoder(w).Encode("Departamento no existente")
	}
}

//Obtiene tienda mediante parametros departamento, nombre y calificacion
func getTiendaEspecifica(w http.ResponseWriter, r *http.Request){
	var aux int
	var tienda ListaDoblementeEnlazada.TiendaIntroducida
	Tienda, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Datos no validos")
	}
	json.Unmarshal([]byte(Tienda), &tienda)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	key := AdminJSON.KeyDepto(departamentos, tienda.Departamento)
	if key != -1{
		aux = getKeyLetra(tienda.Nombre)
		if aux != -1{
			json.NewEncoder(w).Encode(lista[key+len(departamentos)*(aux+27*(tienda.Calificacion-1))].Buscar(tienda))
		}
	}else{
		json.NewEncoder(w).Encode("Departamento no existente")
	}
}

//Generar JSON
func saveJSON(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	AdminJSON.EncoderJson(db)
	json.NewEncoder(w).Encode("Datos Guardados")
}

func main() {
	router := mux.NewRouter()
	//Funcional
	router.HandleFunc("/cargartienda", setJSON).Methods("POST")
	//Funcional
	router.HandleFunc("/Tiendas", getJSON).Methods("GET")
	//Funcional
	router.HandleFunc("/Eliminar", eliminar).Methods("DELETE")
	//Funcional
	router.HandleFunc("/TiendaEspecifica", getTiendaEspecifica).Methods("POST")
	//Funcional
	router.HandleFunc("/guardar", saveJSON).Methods("GET")
	//no Funcional
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}

//Obtencion de llave de letra
func getKeyLetra(nombre string) int {
	var num int = int(nombre[0])
	var runes []rune
	runes = []rune(nombre)
	if (num >= 65 && num <= 78) || (num >= 97 && num <= 110){
		if num >= 65 && num <= 78{
			return num -65
		}else{
			return num -97
		}
	}else if string(runes[0]) == "ñ" || string(runes[0]) == "Ñ"{
		return 14
	}else if (num >= 79 && num <= 90) || (num >= 111 && num <= 122){
		if num >= 79 && num <= 90{
			return num -64
		}else{
			return num -96
		}
	}
	return -1
}


func del(tienda ListaDoblementeEnlazada.TiendaEliminar){
	var ListaVacia []ListaDoblementeEnlazada.Tienda
	var Encontrada bool
	var index int
	var departamento string
	for i:=0; i<len(db.Datos);i++{
		for j:=0; j<len(db.Datos[i].Departamentos);j++{
			for key, value := range db.Datos[i].Departamentos[j]{
				for a, valu := range value{
					if valu.Nombre == tienda.Nombre && valu.Calificacion == tienda.Calificacion && key == tienda.Categoria{
						Encontrada = true
						index = a
						departamento = key
					}
				}
			}
			if Encontrada {
				if len(db.Datos[i].Departamentos[j][departamento]) == 1{
					db.Datos[i].Departamentos[j][departamento] = ListaVacia

				}else{
					db.Datos[i].Departamentos[j][departamento] = append(db.Datos[i].Departamentos[j][departamento][:index], db.Datos[i].Departamentos[j][departamento][index+1:]...)
				}
				return
			}
		}
	}
}

/*func reLinealizar(){
	//var base AdminJSON.DB_VirtualMall
	//var Indice string
	//var runes []rune
	var mapaIndices map[(string)int]
	MatrizTiendas := make([][][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, len(departamentos))
	for i := range MatrizTiendas {
		MatrizTiendas[i] = make([][]ListaDoblementeEnlazada.ListaDoblementeEnlazada, 27)
	}
	for i := range MatrizTiendas {
		for j:= range MatrizTiendas[i]{
			MatrizTiendas[i][j] = make([]ListaDoblementeEnlazada.ListaDoblementeEnlazada, 5)
		}
	}
	for i:=0; i < len(departamentos); i++{
		for j:=0; j < 27; j++{
			for k:=0; k < 5; k++{
				MatrizTiendas[i][j][k] = lista[i+len(departamentos)*(j+27*k)]
			}
		}
	}

	for i:=0; i < len(departamentos); i++{
		for j:=0; j < 27; j++{
			for k:=0; k < 5; k++{
				db.Datos[i].i
				MatrizTiendas[i][j][k] = lista[i+len(departamentos)*(j+27*k)]
			}
		}
	}

}*/
