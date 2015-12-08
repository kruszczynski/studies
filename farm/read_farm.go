package farm

import "fmt"

func readFarm() []*animal {
	animals := []*animal{}
	for shouldContinue := "y"; shouldContinue == "y"; {
		// Name
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
		animals = append(animals, animal)
		// continue
		fmt.Print("Type 'y' to continue, anything else to exit: ")
		fmt.Scan(&shouldContinue)
	}
	return animals
}
