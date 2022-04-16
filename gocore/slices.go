package gocore

import (
	"fmt"
	"strings"
)

func SlicesExample() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	// First index included; last index excluded
	terrestrial := planets[0:4] // [:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8] // [6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	terrestrial[0] = "Mercury"
	fmt.Println(planets)
	allPlanets := planets[:] // full slice
	fmt.Println(allPlanets)
	// Array slices modify initial array
}

func SlicesForStrings() {
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)
	neptune = "Poseidon"
	fmt.Println(tune)
	// String slices not modify initial string
}

func SlicesIndexFeature() {
	question := "¿Cómo estás?"
	fmt.Println(question[:6]) // must be first 7 symbols
	// indexes considering bytes not runes
}

// Composits literals
// [...]{"", "", ""} - array
// []{"", "", ""} - slices

func hyperspace(worlds []string) { // this arg is slice
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

func SliceExampleTrim() {
	planets := []string{" Venera ", " Earth ", " Mars "}
	hyperspace(planets)
	fmt.Println(planets)
	fmt.Println(strings.Join(planets, ""))
}

// Slices with methods
// broken lesson

//type StringSlice []string

//func (p StringSlice) Sort()

//func SliceWithMethods() {
//	planets := []string{
//		"Mercury", "Venera", "Earth", "Mars",
//		"Jupiter", "Saturn", "Uran", "Neptun",
//	}
//	sort.StringSlice(planets).Sort()
//	fmt.Println(planets)
//}

// append and make

func SlicesAppendExample() {
	dwarfs := []string{"Cecera", "Pluton", "Haumea", "Mamake", "Erida"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Salashia", "Kvazar", "Sedna") // multiple append
	fmt.Println(dwarfs)
}

//  length and capacity

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func SlicesLengthExample() {
	dwarfs := []string{"Cecera", "Pluton", "Haumea", "Mamake", "Erida"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
}

//

func SlicesThreeIndex() {
	planets := []string{
		"Mercury", "Venera", "Earth", "Mars",
		"Jupiter", "Saturn", "Uran", "Neptun",
	}
	terrestrial := planets[0:4:4] // len=4 cap=4
	// if cap not setted cap will be 8
	worlds := append(terrestrial, "Cerera")
	worlds = append(terrestrial, "Yaju")
	fmt.Println(planets)
	fmt.Println(worlds)
	// if you dont need rewrite base array you need use three index slice
}

// make
// using make provide more perf
// not future copy and move
func SlicesMakeExample() {
	dwarfs := make([]string, 0, 10) // len 0 cap 10
	dwarfs = append(dwarfs, "Cecera", "Pluton", "Haumea", "Mamake")
	fmt.Println(dwarfs)
	dump("dwarfs", dwarfs)
}

// Variative functions
func terraform(prefix string, worlds ...string) []string {
	newWorld := make([]string, len(worlds)) // newWorld prevent modifying
	for i := range worlds {
		newWorld[i] = prefix + " " + worlds[i]
	}
	return newWorld
}

func SlicesVarFuncs() {
	twoWorld := terraform("New", "Venera", "Mars")
	fmt.Println(twoWorld)
	planets := []string{"Venera", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}
