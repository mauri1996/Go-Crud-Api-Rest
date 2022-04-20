package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// encoding/json -> Codificacion y decodificacion en Json
// log -> logs
// net/http -> Generacion de servidor, peticiones, respuestas
// gihub.com/gorilla/mux ->

// Estructura de datos
type Person struct {
	ID        string   `json:"id,omitempy"`
	FirstName string   `json:"firstname,omitempy"`
	LastName  string   `json:"lastname,omitempy"`
	Address   *Address `json:"address,omitempy"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// Arreglo de personas

var people []Person

// r *http.Request -> obtiene toda la informacion enviada
// w http.ResponseWriter -> almacena la informacion a enviar

// Funcion que devolvera informacion
func GetPeopleEndPoint(w http.ResponseWriter, r *http.Request) { // los parametros son indispensables
	json.NewEncoder(w).Encode(people) // codifica en formato Json, y envia people al servidor- revisar apiTest.rest (1)

}

func GetPersonEndPoint(w http.ResponseWriter, r *http.Request) { // los parametros son indispensables
	params := mux.Vars(r) //obtener informacion
	for _, item := range people {
		if item.ID == params["id"] { // obtener parametro "ID"
			json.NewEncoder(w).Encode(item) // envia dato  // codifica en formato Json, y envia people al servidor- revisar apiTest.rest (2)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{}) // devuelve vacio
}

func CreatePersonEndPoint(w http.ResponseWriter, r *http.Request) { // los parametros son indispensables
	params := mux.Vars(r) //obtener informacion de id
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person) // Decodificar y adignar a la estructura de Persona
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people) // codifica en formato Json, y envia people al servidor- revisar apiTest.rest (3)
}

func DeletePersonEndPoint(w http.ResponseWriter, r *http.Request) { // los parametros son indispensables
	params := mux.Vars(r) //obtener informacion
	for index, item := range people {
		if item.ID == params["id"] { // obtener parametro "ID"
			people = RemoveIndex(people, index)
			json.NewEncoder(w).Encode(people) // envia dato  // codifica en formato Json, y envia people al servidor- revisar apiTest.rest (4)
			return
		}
	}
	json.NewEncoder(w).Encode(people) // devuelve todos los datos
}

// Funcion creada para eliminar datos de un Slice
func RemoveIndex(s []Person, index int) []Person {
	return append(s[:index], s[index+1:]...)
}

func main() {
	// Creacion de enrutador
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Mauri", LastName: "C", Address: &Address{City: "cuidad", State: "stateTest"}})
	people = append(people, Person{ID: "2", FirstName: "aaa", LastName: "B", Address: &Address{City: "a12", State: "fff"}})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	// Crear servidor
	// http.ListenAndServe(":8080", router)

	log.Fatal(http.ListenAndServe(":3000", router)) // para manejar el error si pasa algo
}
