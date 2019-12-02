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

	fmt.Println("\t1202 program alarm:\t", result[0])

	gravityAssist, err = inputToInt(inputs)
	checkErr(err, "unable to convert")
	fmt.Println("\t19690720 output:\t", day2B(inputs))

}

func day2B (inputs []string) int {
	gravityAssist, _ := inputToInt(inputs)
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			gravityAssist[1] = i
			gravityAssist[2] = j

			result, err := intcodeComputer(gravityAssist)
			checkErr(err, "failed input combination")

			if result[0] == 19690720 {
				return 100 * i + j
			}
			gravityAssist, _ = inputToInt(inputs)
		}
	}
	return 0
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
