package main

import (
	"encoding/json"
	"net/http"
	"io"

	"github.com/julienschmidt/httprouter"
)

var animals = []*animal{}

func serializeAnimals(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(animals)
}

// GetAnimals renders animals list
func getAnimals(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	serializeAnimals(w)
}

func createAnimal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	animal := &animal{}
	json.NewDecoder(r.Body).Decode(&animal)
	defer r.Body.Close()
	animals = append(animals, animal)
}

func main() {
	router := httprouter.New()
	router.GET("/animals", getAnimals)
	router.POST("/animals", createAnimal)
	http.ListenAndServe(":8080", router)
}
