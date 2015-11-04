package main

import (
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

var animals = []*animal{}

func serializeAnimals(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(animals); err != nil {
		return err
	}
	return nil
}

// GetAnimals renders animals list
func getAnimals(w http.ResponseWriter, _r *http.Request, _ httprouter.Params) {
	serializeAnimals(w)
}

func createAnimal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	animal := &animal{}
	// animal.Name = r.FormValue("name")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
	json.Unmarshal(body, &animal)
	animals = append(animals, animal)
}

func main() {
	router := httprouter.New()
	router.GET("/animals", getAnimals)
	router.POST("/animals", createAnimal)
	http.ListenAndServe(":8080", router)
}
