package main

import "fmt"
import "encoding/json"
import "io/ioutil"

type Species int

const (
  Cow Species = iota
  Pig
  Cock
  Tit
  Deer
  Sheep
)

var SpeciesNames = [6]string {
  "Cow",
  "Pig",
  "Cock",
  "Tit",
  "Deer",
  "Sheep",
}

// serialize enum values
func (species Species) MarshalJSON() ([]byte, error) {
  return json.Marshal(SpeciesNames[species])
}

type Animal struct {
  Species Species
  Name string
  Age uint
}

type Farm struct {
  Name string
  Animals []Animal
}

func readData() Farm {
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
    var newSpecies Species
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

func serializeData(farm Farm) {
  farmBytes, serializeError := json.Marshal(farm)
  fmt.Println(string(farmBytes))
  if serializeError != nil {
    panic(serializeError)
  }
  writeError := ioutil.WriteFile("farm.json", farmBytes, 0644)
  if writeError != nil {
    panic(writeError)
  }
  fmt.Println("Farm data saved successfully to farm.json")
}

func main() {
  fmt.Println("Welcome to place sunlight never reaches...")
  farm := readData()
  fmt.Println("Farm:", farm.Name)
  serializeData(farm)
}