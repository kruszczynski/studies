package main

import (
	"fmt"
	"encoding/json"
	"os"
	"net/http"
)

type species int

const (
	cow species = iota
	pig
	cock
	tit
	deer
	sheep
)

const fileName string = "farm.json"

var speciesNames = [6]string{
	"Cow",
	"Pig",
	"Cock",
	"Tit",
	"Deer",
	"Sheep",
}

// MarshalJSON intr
func (species species) MarshalJSON() ([]byte, error) {
	return json.Marshal(speciesNames[species])
}

// Animal is a representation of an animal
type animal struct {
	Species species
	Name    string
	Age     uint
}

func readFarm() []*animal {
	animals := []*animal{}
	for shouldContinue := "y"; shouldContinue == "y"; {
		// name
		animal := &animal{}
		fmt.Print("Animal's name: ")
		fmt.Scan(&animal.Name)
		// species
		fmt.Print("Animal's species (0=Cow,1=Pig,2=Cock,3=Tit,4=Deer,5=Sheep): ")
		fmt.Scanf("%d", &animal.Species)
		animal.Species = animal.Species % species(len(speciesNames))
		// age
		fmt.Print("Animal's age: ")
		fmt.Scanf("%d", &animal.Age)
		animals = append(animals, animal)
		// continue
		fmt.Print("Type 'y' to continue, anything else to exit: ")
		fmt.Scan(&shouldContinue)
	}
	return animals
}

func serializeFarm(animals []*animal) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(animals); err != nil {
		return err
	}
	fmt.Println("Farm data saved successfully to", fileName)
	return nil
}

func main() {
	if err := serializeFarm(readFarm()); err != nil {
		panic(err)
	}
  // http.HandleFunc("/", handler)
  // http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
