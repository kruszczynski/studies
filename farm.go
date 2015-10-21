package main

import "fmt"
import "encoding/json"
// import "io/ioutil"
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

func readFarm() *Farm {
  fmt.Print("What is your farm's name: ")
  var farmName string
  fmt.Scan(&farmName)
  farm := &Farm{Name: farmName}
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
    newSpecies = newSpecies % 6
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

func serializeFarm(farm *Farm) error {
  file, err := os.Create(fileName)
  if err != nil {
    return(err)
  }
  encoder := json.NewEncoder(file)
  if err := encoder.Encode(farm); err != nil {
    return(err)
  }
  fmt.Println("Farm data saved successfully to", fileName)
  return(nil)
}

func main() {
  fmt.Println("Welcome to place sunlight never reaches...")
  farm := readFarm()
  if err := serializeFarm(farm); err != nil {
    panic(err)
  }
}