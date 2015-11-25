package main

import (
	"encoding/json"
	"fmt"
)

type species int

// Animal is a representation of an animal
type animal struct {
	Species species
	Name    string
	Age     uint
}

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

func (a *animal) UnmarshalJSON(data []byte) error {
	tempAnimal := &animal{}
	if err := json.Unmarshal(data, &tempAnimal); err != nil {
		return err
	}
	fmt.Print(tempAnimal)
	return nil
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
	return fmt.Errorf("no such animal: %s", entry)
}
