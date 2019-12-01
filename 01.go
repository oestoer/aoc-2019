package main

import (
	"fmt"
	"math"
)

// Day1A returns the solution for the first part of 'The Tyranny of the Rocket Equation'
func Day1A(moduleMasses []int, debug bool) int {
	var fuelReq []int
	var sum int

	for i := 0; i < len(moduleMasses); i++ {
		fuelReq[i] = fuel(moduleMasses[i])
		sum += fuelReq[i]
	}
	if debug {
		fmt.Println(sum, fuelReq)
	}

	return sum
}

// Day1B returns the solution for the first part of 'The Tyranny of the Rocket Equation'
func Day1B(moduleMasses []int, debug bool) int {
	var fuelReq []int
	var sum int

	for i := 0; i < len(moduleMasses); i++ {
		fuelReq[i] = fuelWithAddedFuel(moduleMasses[i])
		sum += fuelReq[i]
	}
	if debug {
		fmt.Println(sum, fuelReq)
	}

	return sum
}

func fuel(m int) int {
	return int(math.Floor(float64(m/3))) - 2
}

func fuelWithAddedFuel(f int) int {
	r := fuel(f)
	if r <= 0 {
		return 0
	}
	r += fuelWithAddedFuel(r)

	return r
}
