package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// randomPeck returns a boolean for a random peck:
// true maps to a peck to the right, false to the left.
func randomPeck() bool {
	return rand.Float32() < 0.5
}

func runSim() int {
	var chicks [100]bool
	var isPecked [100]bool

	for i := range chicks {
		chicks[i] = randomPeck()
	}

	for i, chick := range chicks {
		if chick {
			var tgt int
			if i < 99 {
				tgt = i + 1
			} else {
				tgt = 0
			}
			isPecked[tgt] = true
		} else {
			var tgt int
			if i > 0 {
				tgt = i - 1
			} else {
				tgt = 99
			}
			isPecked[tgt] = true
		}
	}

	var totalPecked int
	for _, pecked := range isPecked {
		if pecked {
			totalPecked++
		}
	}

	return totalPecked
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sims := 1000000
	var total int
	for i := 0; i < sims; i++ {
		total += runSim()
	}
	fmt.Printf("After a total of %d simulations, the average pecked chicks are %f", sims, math.Round(float64(total/sims)))
}
