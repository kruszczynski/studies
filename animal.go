package main

import "encoding/json"

type species int

const (
  cow species = iota
  pig
  cock
  tit
  deer
  sheep
)

var speciesNames = [6]string{
  "Cow",
  "Pig",
  "Cock",
  "Tit",
  "Deer",
  "Sheep",
}

// MarshalJSON intr
func (s species) MarshalJSON() ([]byte, error) {
  return json.Marshal(speciesNames[s])
}

func (s species) UnmarshalJSON(byte []byte) error {
  for index, name := range speciesNames {
    trimmedInput := string(byte[1:len(byte)-1])
    if name == trimmedInput {
      s = species(index)
    }
  }
  return nil
}

// Animal is a representation of an animal
type animal struct {
  Species species
  Name    string
  Age     uint
}