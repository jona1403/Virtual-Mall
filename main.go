package main

import (
	"encoding/json"
	"fmt"
	"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"net/http"
	//"github.com/VirtualMall/ListaDoblementeEnlazada/AdminJSON"
)
var db AdminJSON.DB_VirtualMall

func getJSON(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}

func setJSON(w http.ResponseWriter, r *http.Request){
	Tiendas, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Datos no validos")
	}
	json.Unmarshal([]byte(Tiendas), &db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	AdminJSON.EncoderJson(db)
	json.NewEncoder(w).Encode("Datos cargados")
}

func saveJSON(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Datos Guardados")
}

func getArreglo(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Arreglo creado exitosamente")
}

func Eliminar(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Tienda eliminada")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cargartienda", setJSON).Methods("POST")
	router.HandleFunc("/Tiendas", getJSON).Methods("GET")
	router.HandleFunc("/guardar", saveJSON).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}