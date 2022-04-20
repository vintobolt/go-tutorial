package gocore

import (
	"fmt"
	"sort"
)

// not initiated types
// int -> 0
// string -> ""
// pointer -> nil
// maps, slices, interfaces -> nil

func NilPointerErrExample() {
	var nowhere *int
	fmt.Println(nowhere)
	fmt.Println(*nowhere)
}

func NilPointerErrFixExample() {
	var nowhere *int
	if nowhere != nil {
		fmt.Println(nowhere)
		fmt.Println(*nowhere)
	}

}

// Securing methods
type person struct {
	age int
}

func (p *person) birthday() {
	p.age++ // err here
}

// seccuring method
func (p *person) birthdayFix() {
	if p == nil {
		return
	}
	p.age++
}

func NilPointerStructErr() {
	var nobody *person // pointer to nil
	fmt.Println(nobody)
	nobody.birthdayFix()
}

// nil types of funcs

var fn func(a int, b int) int

func NilFuncsType() {
	fmt.Println("fn == nil: ", fn == nil)
}

// if function nil
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less) // default if nil
}

func NilFunctionSortExample() {
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)
}

// nil slices

func NilSlices() {
	// Slice created without composit literal or make func will be nil
	var soup []string
	fmt.Println("soup is nil: ", soup == nil)
	for _, ingridient := range soup { // but we can use append, len, range...
		fmt.Println(ingridient)
	}
	fmt.Println(len(soup)) // 0
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)
	// but nill is not equal empty slice
}

///
func mirepoix(ingridients []string) []string {
	return append(ingridients, "onion", "carrot", "celery")
}

func NilSlicesExample() {
	soup := mirepoix(nil)
	fmt.Println(soup)
}

// Nil maps

func NilMapsExample() {
	// announcement
	var soup map[string]int
	fmt.Println(soup == nil) // true
	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingridient, measmeasurement := range soup {
		fmt.Println(ingridient, measmeasurement)
	}
}

// Nil interfaces
var v interface{}

func NilInterfacesPrint() {
	fmt.Printf("%T %v %v\n", v, v, v == nil)
}

func NilInterfaceExample() {
	var p *int
	v = p
	fmt.Printf("%#v\n", v)
}
