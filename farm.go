package main

import "fmt"

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

type Animal struct {
  species Species
  name string
  age uint
}

type Farm struct {
  animals []Animal
}

func main() {
  // a := Animal{species: Cow, name: "Honey", age: 5}
  farm := new(Farm)
  // farm.animals = append(farm.animals, a)
  // ani := farm.animals[0]

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
    newAnimal := Animal{name: newName, species: newSpecies, age: uint(newAge)}
    farm.animals = append(farm.animals, newAnimal)
    // continue
    fmt.Print("Type 'y' to continue, anything else to exit: ")
    fmt.Scan(&shouldContinue)
  }
  for _, animal := range farm.animals {
    fmt.Println("Next Animal")
    fmt.Println(animal.name)
    fmt.Println(animal.age)
    fmt.Println(SpeciesNames[animal.species])
  }
}