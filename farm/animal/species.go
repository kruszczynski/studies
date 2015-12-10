package animal

import (
	"encoding/json"
	"fmt"
)

// Species is Animal's species
type Species int

const (
	cow Species = iota
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

// UnmarshalJSON Deserializes to Species from string
func (s Species) UnmarshalJSON(data []byte) error {
	var entry string
	fmt.Println(entry)
	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}
	fmt.Println(entry)
	for i, name := range speciesNames {
		if name == entry {
			// for whatever reason that does not persist
			// therefore its all cows
			s = Species(i)
			return nil
		}
	}
	return fmt.Errorf("no such animal: %s", entry)
}

// MarshalJSON intr
func (s Species) MarshalJSON() ([]byte, error) {
	return json.Marshal(speciesNames[s])
}
