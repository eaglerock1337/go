package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	totalSampleSize = 1000000
	diceSets        = 3
)

/*
Roll and attempt to hit at least one six per set of 6 dice.
Returns true if at least dice_sets sixes are rolled. If ceiling
is set to True, will only return True for an exact hit.

Variables:
  numOfSets - The number of dice to roll (in sets of 6)
  ceiling   - Whether an exact hit is required or if rolling
              more sixes still counts as a hit.

Returns a True result if the number of sixes matches
the number of sets, based on the ceiling strategy.
*/
func rollForSixes(numOfSets int, ceiling bool) bool {
	totalDice := numOfSets * 6
	sixes := 0

	for i := 0; i < totalDice; i++ {
		roll := rand.Intn(6) + 1
		if roll == 6 {
			sixes++
		}
	}

	if ceiling {
		return (sixes == numOfSets)
	}
	return (sixes >= numOfSets)
}

/*
Runs a simulation of the specified sample size and roll
strategy. Prints out data and results of the simulation
to stdout.

Variables:
  sample_size - how many times we test each chest
  exact_hit   - whether an exact number of sixes is required
				for a hit, or if rolling extra sixes will
				still count.
*/
func runSimulation(sampleSize int, exactHit bool) {
	var successes [diceSets]int

	for i := 0; i < sampleSize; i++ {
		for sets := 1; sets <= diceSets; sets++ {
			result := rollForSixes(sets, exactHit)
			if result {
				successes[sets-1]++
			}
		}
	}
	fmt.Printf("\n\nResults for sample size %v (Exact hit: %v):\n", sampleSize, exactHit)
	for i := 1; i <= diceSets; i++ {
		fmt.Printf("\nUsing %v sets of dice (%v total):\n", i, i*6)
		fmt.Printf("  Total Hits: %v\n", successes[i-1])
		fmt.Printf("  Winning Percentage: %.6v%%\n", float64(successes[i-1])/float64(sampleSize)*100.0)
	}
}

// Simulates the Newton–Pepys problem
// Inspired by https://www.youtube.com/watch?v=RFlTawWwLZc
func main() {
	rand.Seed(time.Now().UnixNano())
	runSimulation(totalSampleSize, false)
	runSimulation(totalSampleSize, true)
}
