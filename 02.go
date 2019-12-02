package main

import (
	"fmt"
)

func Day2() {
	inputs, err := readInput(2, ",")
	checkErr(err, "unable to read file")

	gravityAssist, err := inputToInt(inputs)
	checkErr(err, "unable to convert")

	gravityAssist[1] = 12
	gravityAssist[2] = 2
	result, err := intcodeComputer(gravityAssist)
	checkErr(err, "1202")

	fmt.Println("\t1202 program alarm: ", result[0])

}

func intcodeComputer(input []int) ([]int, error) {
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			input[input[i+3]] = input[input[i+2]] + input[input[i+1]]
		case 2:
			input[input[i+3]] = input[input[i+2]] * input[input[i+1]]
		case 99:
			return input, nil
		default:
			return nil, fmt.Errorf("unknown instruction %v", input[i])
		}

	}

	return input, nil
}
