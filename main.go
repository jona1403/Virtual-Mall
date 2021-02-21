package main

import (
	"encoding/json"
	"fmt"
	"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
	"github.com/VirtualMall/ListaDoblementeEnlazada/ListaDoblementeEnlazada"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
)
var db AdminJSON.DB_VirtualMall
var lista []ListaDoblementeEnlazada.ListaDoblementeEnlazada
var departamentos map[int]string
var indices map[int]string

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
	lista, departamentos, indices = AdminJSON.Linealizacion(db)

	json.NewEncoder(w).Encode("Datos cargados")
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
		aux = AdminJSON.KeyIndice(indices, tienda.Nombre)
		if aux != -1{
			del(tienda)
			lista[key+len(departamentos)*(aux+len(indices)*(tienda.Calificacion-1))].Eliminar(tienda.Nombre)
			json.NewEncoder(w).Encode("Ok")
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
		aux = AdminJSON.KeyIndice(indices, tienda.Nombre)
		if aux != -1{
			json.NewEncoder(w).Encode(lista[key+len(departamentos)*(aux+len(indices)*(tienda.Calificacion-1))].Buscar(tienda))
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

//Obtencion de arreglo grafico
func getArreglo(w http.ResponseWriter, r *http.Request){
	AdminJSON.Graficar(lista)
	json.NewEncoder(w).Encode("Grafica generada")
}
//Obtiene posicion especifica
func getPosition(w http.ResponseWriter, r *http.Request){
	datos := mux.Vars(r)
	Posision, _ := strconv.Atoi(datos["id"])
	lpos, _ := strconv.Atoi(datos["numero"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if Posision<len(lista) {
		if lpos <= lista[Posision].Tamano && lpos > 0{
			cabeza := lista[Posision].Cabeza
			for i:= 1; i<=lpos;i++{
				if i == lpos{
					json.NewEncoder(w).Encode(cabeza.Dato)
				}else{
					cabeza = cabeza.Siguiente
				}
			}
		}else{
			json.NewEncoder(w).Encode("Numero invalido")
		}
	}else {
		json.NewEncoder(w).Encode("Posicion invalida")
	}

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
	//Funcional
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")

	router.HandleFunc("/{id}/{numero}", getPosition).Methods("GET")
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
