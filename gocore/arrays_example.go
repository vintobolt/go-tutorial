package gocore

import "fmt"

func ArrayExample() {
	var planets [8]string
	planets[0] = "Меркурий"
	planets[1] = "Венера"
	planets[2] = "Земля"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")
}

func ArraysCreation() {
	dwarfs := [5]string{"Венера", "Плутон", "Хаумеа", "Марс", "Эрида"}
	fmt.Println(dwarfs)
	planets := [...]string{
		"Венера",
		"Плутон",
		"Марс",
		"Сатурн",
	}
	fmt.Println(planets)
}

func ArraysIteration() {
	dwarfs := [5]string{"Венера", "Плутон", "Хаумеа", "Марс", "Эрида"}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
