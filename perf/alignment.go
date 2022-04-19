package perf

import (
	"fmt"
	"unsafe"
)

type someStruct struct {
	l          bool
	m          bool
	someField  int32
	someField1 int32
	//l          bool
}

func PerfMemoryAlign() {
	sm := someStruct{someField: 12, someField1: 12, l: false, m: true}
	fmt.Println(unsafe.Sizeof(sm))
}

type Foo struct {
	aaa [2]bool
	bbb int32
	ccc [2]bool
}

type Bar struct {
	aaa [2]bool
	bbb [2]bool
	ccc int64
}

func PerfAlignOf() {
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(Foo{}))
	fmt.Printf("required alignment: %+v\n", unsafe.Alignof(Bar{}))
}
