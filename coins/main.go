package main

import (
	"fmt"
	"math/rand"
	"time"
)

const sampleSize = 1000000000

/*
Runs a coin flip simulation for the specified amount of flips.
Tallies and prints out relevant data for the run.
*/
func runSimulation(flips int) {
	var sides = [2]string{"Heads", "Tails"}
	var totalFlips, totalStreaks [2]int // The total number of flips and streaks per side
	var streaks []int                   // Starts as a slice of length 1 and gets extended as streaks grow

	curStreak, maxStreak, lastFlip := 1, 0, -1

	// Main simulation loop
	fmt.Printf("\nSimulating %v coin flips...\n\n", flips)
	for i := 0; i < flips; i++ {
		flip := rand.Intn(2)
		totalFlips[flip]++

		if flip == lastFlip {
			curStreak++
		}

		if flip != lastFlip || i == flips-1 {
			if maxStreak < curStreak {
				for i := len(streaks); i < curStreak; i++ {
					streaks = append(streaks, 0)
				}
				maxStreak = curStreak
			}
			totalStreaks[flip]++
			streaks[curStreak-1]++
			curStreak = 1
			lastFlip = flip
		}
	}

	// Print out fancy-schmancy results
	for i := 0; i <= 1; i++ {
		percent := float64(totalFlips[i]) / float64(flips) * 100.0
		fmt.Printf("%v:  Total Flips: %6v  Total Streaks: %5v  Percentage: %.6v%%\n", sides[i], totalFlips[i], totalStreaks[i], percent)
	}
	fmt.Printf("\nThe biggest streak was %v.\n", maxStreak)
	fmt.Println("Streak counts:")
	for i := 0; i < len(streaks); i++ {
		if streaks[i] > 0 {
			fmt.Printf("%5v: %10v\n", i+1, streaks[i])
		}
	}
}

// Simulates coin flips to determine how many and how long
// streaks of either heads or tails will occur.
// Inspired by https://www.youtube.com/watch?v=DntEoGG7RyY
func main() {
	rand.Seed(time.Now().UnixNano())
	runSimulation(sampleSize)
}
