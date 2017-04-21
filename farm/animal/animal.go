package animal

// Animal is a representation of an animal
type Animal struct {
	Species Species `json:"Species"`
	Name    string  `json:"Name"`
	Age     uint    `json:"Age"`
}
