package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const kimBusy = true

// Bell is a thing.
type Bell struct {
	ringing bool
	mux     sync.Mutex
}

// Student is a thing.
type Student struct {
	name      string
	busyCount int
	response  chan string
}

func newStudent(name string) Student {
	s := Student{name: name}
	s.busyCount = 0
	s.response = make(chan string)
	return s
}

func kim(hand chan Student) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("Kim's all ready to help!")

	for {
		somebody := <-hand
		fmt.Printf("Kim sees %v's raised hand!\n", somebody.name)
		if kimBusy {
			if somebody.busyCount < 5 {
				fmt.Println("Kim says 'Busy!'")
				somebody.response <- "Busy"
			} else {
				fmt.Println("Kim says 'You're on Deck!'")
				somebody.response <- "On Deck"
			}
		} else {
			fmt.Println("Something went wrong! Kim is always busy!")
		}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func student(wg *sync.WaitGroup, name string, hand chan Student) {
	defer wg.Done()
	me := newStudent(name)
	fmt.Printf("%v has started his work!\n", me.name)
	time.Sleep(time.Duration(rand.Intn(5)+5) * time.Second)
	fmt.Printf("%v needs some help and raises his hand!\n", me.name)

	for {
		select {
		case hand <- me:
			kimSez := <-me.response
			if kimSez == "Busy" {
				me.busyCount++
			} else if kimSez == "On Deck" {
				fmt.Printf("%v is on deck and accepts his fate.\n", me.name)
				return
			} else {
				fmt.Println("This will never happen! Kim's always busy!")
				return
			}
			fmt.Printf("%v waits a bit to raise his hand again.\n", me.name)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("%v raises his hand again!\n", me.name)
		default:
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}
}

func classroom(class []string) {
	var wg sync.WaitGroup
	hand := make(chan Student, 1)
	fmt.Println("Time for class!")

	go kim(hand)
	for _, name := range class {
		go student(&wg, name, hand)
	}

	wg.Add(len(class))
	wg.Wait()
	fmt.Println("The bell has rung!")
}

func main() {
	names := []string{
		"Jon", "Peter", "Mike", "Daniel",
		"Ben", "Jared", "Nima", "Mark",
		"Bill", "Ted", "Marty", "Biff",
	}
	rand.Seed(time.Now().UnixNano())
	classroom(names)
}
