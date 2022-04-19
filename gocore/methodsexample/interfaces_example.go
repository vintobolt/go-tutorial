package methodsexample

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// интерфейсы могут использоваться со встраиванием структуры

type starship struct {
	laser
}

func InterfaceExample() {
	shout(martian{})
	shout(laser(2))
	s := starship{laser(3)}
	//fmt.Println(s.talk())
	shout(s)
}
