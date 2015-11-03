package main

import "fmt"
import "encoding/json"
import "os"

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

// Farm is a representation of a Farm that has animals
type Farm struct {
	Name    string
	Animals []*animal
}

func readFarm() *Farm {
	fmt.Print("What is your farm's name: ")
	var farmName string
	fmt.Scan(&farmName)
	farm := &Farm{Name: farmName}
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
		farm.Animals = append(farm.Animals, animal)
		// continue
		fmt.Print("Type 'y' to continue, anything else to exit: ")
		fmt.Scan(&shouldContinue)
	}
	return farm
}

func serializeFarm(farm *Farm) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(farm); err != nil {
		return err
	}
	fmt.Println("Farm data saved successfully to", fileName)
	return nil
}

func main() {
	fmt.Println("Welcome to place sunlight never reaches...")
	if err := serializeFarm(readFarm()); err != nil {
		panic(err)
	}
}
