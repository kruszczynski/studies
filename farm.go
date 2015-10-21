package main

import "fmt"
import "encoding/json"
import "io/ioutil"

type species int

const (
  cow species = iota
  pig
  cock
  tit
  deer
  sheep
)

var speciesNames = [6]string {
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
type Animal struct {
  Species species
  Name string
  Age uint
}

// Farm is a representation of a Farm that has animals
type Farm struct {
  Name string
  Animals []Animal
}

func readFarm() Farm {
  fmt.Print("What is your farm's name: ")
  var farmName string
  fmt.Scan(&farmName)
  farm := Farm{Name: farmName}
  shouldContinue := "y"
  for shouldContinue == "y" {
    // name
    fmt.Print("Animal's name: ")
    var newName string
    fmt.Scan(&newName)
    // species
    fmt.Print("Animal's species (0=Cow,1=Pig,2=Cock,3=Tit,4=Deer,5=Sheep): ")
    var newSpecies species
    fmt.Scanf("%d", &newSpecies)
    // age
    fmt.Print("Animal's age: ")
    var newAge int
    fmt.Scanf("%d", &newAge)
    newAnimal := Animal{Name: newName, Species: newSpecies, Age: uint(newAge)}
    farm.Animals = append(farm.Animals, newAnimal)
    // continue
    fmt.Print("Type 'y' to continue, anything else to exit: ")
    fmt.Scan(&shouldContinue)
  }
  return farm
}

func serializeFarm(farm Farm) {
  farmBytes, err := json.Marshal(farm)
  if err != nil {
    panic(err)
  }
  err = ioutil.WriteFile("farm.json", farmBytes, 0644)
  if err != nil {
    panic(err)
  }
  fmt.Println("Farm data saved successfully to farm.json")
}

func main() {
  fmt.Println("Welcome to place sunlight never reaches...")
  farm := readFarm()
  serializeFarm(farm)
}