package gocore

import "fmt"

func PointersTypeExample() {
	a := 5
	b := "Test"
	fmt.Printf("Pointer type is: %T\n", &a)
	fmt.Printf("Pointer type is: %T\n", &b)
}

// struct
type SomeStruct struct {
	a int
}

func (s SomeStruct) AddWithoutPointer(summ int) {
	s.a += summ
}

func (s *SomeStruct) Add(summ int) {
	s.a += summ
}

func PointerFuncWithPointerArgs(s *SomeStruct) {
	s.a = 12
}

func PointersStructureWithout() {
	someStruct := &SomeStruct{a: 15}
	secondStruct := SomeStruct{a: 14}
	fmt.Printf("Somestruct: %+v\n", someStruct)
	someStruct.Add(5)
	fmt.Printf("someStruct: %+v\n", someStruct)
	fmt.Printf("secondStruct: %+v\n", secondStruct)
	PointerFuncWithPointerArgs(&secondStruct)
	fmt.Printf("secondStruct: %+v\n", secondStruct)
}

// maps are pointers already
// Pointers really needs for arrays and structs
