package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println()
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "-h" {
		fmt.Println("Usage:\taoc-2019 COMMAND")
		fmt.Println()
		fmt.Println("Advent of Code 2019 🎄 🕯️ 📅")
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("  all\t\tList all solutions")
		fmt.Println("  day\t\tList solution for given day, where 'day' is a number between 1-25")
		fmt.Println("  help\t\tShow this help")
		os.Exit(0)
	}

	days := map[uint8]func(){
		1: Day1,
		2: Day2,
	}
	if os.Args[1] == "all" {
		for i := 1; i <= len(days); i++ {
			fmt.Printf("📅 Day %d:\n", i)
			days[uint8(i)]()
			fmt.Println("________________________________________________________________________________")
			fmt.Println()
		}
		os.Exit(0)
	}

	d, err := strconv.Atoi(os.Args[1])
	checkErr(err, "not a number")
	day := uint8(d)

	fmt.Printf("📅 Day %d:\n", day)
	days[day]()
}
