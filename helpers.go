package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Printf("%v %v\n", msg, err.Error())
		os.Exit(2)
	}
}

// inputToInt converts a slice of string containing numbers to a slice of int
func inputToInt(err error) func(inputs []string) ([]int, error) {
	return func(inputs []string) ([]int, error) {
		c := make([]int, len(inputs))
		for i, v := range inputs {
			c[i], err = strconv.Atoi(v)
			if err != nil {
				return c, err
			}
		}
		return c, nil
	}
}

// readInput gets the content of an input file from the `inputs` directory for a given day
func readInput(day uint8) ([]string, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("./inputs/day%d.txt", day))
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(file), "\n")
	return s[:len(s)-1], nil
}