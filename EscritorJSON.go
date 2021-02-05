package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion string
}

type Departamento struct {
	Nombre string
	Tiendas []Tienda
}

type Datos_Indices struct {
	Indece string
	Departamentos []Departamento
}
type DB_VirtualMall struct {
	Datos []Datos_Indices
}

func main3(){

	Tiendas := []Tienda{{Nombre: "Aurora", Descripcion: "Es una empresa multinacional " +
		"estadounidense dedicada al diseño, desarrollo, fabricación y comercialización " +
		"de equipamiento deportivo: balones, calzado, ropa, equipo, accesorios y otros " +
		"artículos deportivos", Contacto: "5544-3377", Calificacion: "5"}}

	Departamentos := []Departamento{{Nombre: "Deportes", Tiendas: Tiendas}}
	Datos2 := []Datos_Indices{{Indece: "A", Departamentos: Departamentos}}
	DB := DB_VirtualMall{Datos: Datos2}

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
