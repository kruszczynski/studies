package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kruszczynski/studies/farm/animal"
	"github.com/kruszczynski/studies/farm/checksum"
	"github.com/kruszczynski/studies/farm/secret"
)

var animals = []*animal.Animal{}

func serializeAnimals(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(animals)
}

// GetAnimals renders animals list
func getAnimals(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	serializeAnimals(w)
}

func createAnimal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	animal := &animal.Animal{}
	json.NewDecoder(r.Body).Decode(animal)
	defer r.Body.Close()
	animals = append(animals, animal)
}

func brides(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	cypherWriter := secret.NewWriter()
	if cypherWriter != nil {
		hashWriter := checksum.NewWriter(cypherWriter)
		serializeAnimals(hashWriter)
		hashWriter.PipeSum(w)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/animals", getAnimals)
	router.POST("/animals", createAnimal)
	router.POST("/free_russian_brides", brides)
	http.ListenAndServe(":8080", router)
}
