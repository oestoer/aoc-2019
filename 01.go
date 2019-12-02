package main

import (
	"fmt"
	"math"
	"os"
)

// Day1 returns the solutions for the 'The Tyranny of the Rocket Equation'
func Day1() {
	inputs, err := readInput(1, "\n")
	checkErr(err, "unable to read file")

	moduleMasses, err := inputToInt(inputs)
	checkErr(err, "unable to convert")

	fmt.Printf("\tA: fuel requirement:\t%v\n", day1A(moduleMasses))
	fmt.Printf("\tB: fuel requirement:\t%v\n", day1B(moduleMasses))
}

// day1A returns the solution for the first part of 'The Tyranny of the Rocket Equation'
func day1A(moduleMasses []int) int {
	fuelReq := make([]int, len(moduleMasses))
	var sum int

	for i, m := range moduleMasses {
		fuelReq[i] = fuel(m)
		sum += fuelReq[i]
	}
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(sum, fuelReq)
	}

	return sum
}

// day1B returns the solution for the first part of 'The Tyranny of the Rocket Equation'
func day1B(moduleMasses []int) int {
	fuelReq := make([]int, len(moduleMasses))
	var sum int

	for i, m := range moduleMasses {
		fuelReq[i] = fuelWithAddedFuel(m)
		sum += fuelReq[i]
	}
	if os.Getenv("DEBUG") == "1" {
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
