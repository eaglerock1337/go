package main

import (
	"fmt"
	"sync"
)

const (
	kimBusyMale   = true
	kimBusyFemale = false
)

type marchese struct {
	queue map[string]int
	mux   sync.Mutex
}

type student struct {
	name   string
	onDeck bool
}

func newStudent(name string) *student {
	s := student{name: name}
	s.onDeck = false
	return &s
}

func raiseHand() {

}

func (s student) doWork() {

}

func main() {
	names := [12]string{
		"Jon", "Peter", "Mike", "Daniel",
		"Ben", "Jared", "Nima", "Mark",
		"Bill", "Ted", "Marty", "Biff",
	}
	fmt.Println(names)
}
