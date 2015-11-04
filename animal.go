package main

import (
  "encoding/json"
  "fmt"
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

func (s species) UnmarshalJSON(data []byte) error {
  var entry string
  if err := json.Unmarshal(data, &entry); err != nil {
    return err
  }
  for i, name := range speciesNames {
    if name == entry {
      // for whatever reason that does not persist
      s = species(i)
      return nil
    }
  }
  return fmt.Errorf("no such animal: %q", entry)
}

// Animal is a representation of an animal
type animal struct {
  Species species
  Name    string
  Age     uint
}