package main

import (
	"fmt"
	"sync"
)

const kimBusy = true

type hands struct {
	queue map[string]int
	mux   sync.Mutex
}

type student struct {
	Name       string
	RaisedHand bool
	OnDeck     bool
}

func newStudent(name string) student {
	s := student{Name: name}
	s.OnDeck = false
	return s
}

func (s student) raiseHand() {

}

func kim() {

}

func classroom(class []string) {
	theClass := make([]student, len(class))

	for _, name := range class {
		theClass = append(theClass, newStudent(name))
	}
}

func main() {
	names := []string{
		"Jon", "Peter", "Mike", "Daniel",
		"Ben", "Jared", "Nima", "Mark",
		"Bill", "Ted", "Marty", "Biff",
	}
	go kim()
	go classroom(names)
	fmt.Println(names)
}
