package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x     int
	y     int
	steps int
}

// Day3 returns the solutions for the 'Crossed Wires'
func Day3() {
	inputs, err := readInput(3, "\n")
	checkErr(err, "unable to read file")

	//inputs[0] = "R8,U5,L5,D3"
	//inputs[1] = "U7,R6,D4,L4"

	minMD := shortestManhattanDistance(inputs)
	minSD := shortestSignalDelay(inputs)
	fmt.Println("\tA: closest intersection: ", minMD)
	fmt.Println("\tB: shortest signal delay: ", minSD)
}

func shortestManhattanDistance(inputs []string) int {
	wireCoords := make([][]coord, len(inputs))
	for i, input := range inputs {
		wirePath := readWirePath(input)
		wireCoords[i] = calcWireCoords(wirePath)
	}

	intersections := findIntersectionPoints(wireCoords)
	minD := int(1e10)
	for i := 1; i < len(intersections); i++ {
		m := manhattanDistance(intersections[0], intersections[i])
		if minD > m {
			minD = m
		}
	}
	return minD
}

func shortestSignalDelay(inputs []string) int {
	wireCoords := make([][]coord, len(inputs))
	for i, input := range inputs {
		wirePath := readWirePath(input)
		wireCoords[i] = calcWireCoords(wirePath)
	}

	intersections := findIntersectionPoints(wireCoords)
	minD := int(1e10)
	for i := 1; i < len(intersections); i++ {
		if minD > intersections[i].steps {
			minD = intersections[i].steps
		}
	}
	return minD
}

func readWirePath(wp string) []string {
	return strings.Split(wp, ",")
}

func calcWireCoords(wp []string) []coord {
	wm := []coord{{x: 0, y: 0, steps: 0}}
	var ws int
	for _, w := range wp {
		dir := string(w[0])
		steps, err := strconv.Atoi(w[1:])
		if err != nil {
			continue
		}
		switch dir {
		case "U":
			for j := 1; j <= steps; j++ {
				ws++
				wm = append(wm, coord{
					x:     wm[len(wm)-1].x,
					y:     wm[len(wm)-1].y + 1,
					steps: ws,
				})
			}
		case "D":
			for j := 1; j <= steps; j++ {
				ws++
				wm = append(wm, coord{
					x:     wm[len(wm)-1].x,
					y:     wm[len(wm)-1].y - 1,
					steps: ws,
				})
			}
		case "L":
			for j := 1; j <= steps; j++ {
				ws++
				wm = append(wm, coord{
					x:     wm[len(wm)-1].x - 1,
					y:     wm[len(wm)-1].y,
					steps: ws,
				})
			}
		case "R":
			for j := 1; j <= steps; j++ {
				ws++
				wm = append(wm, coord{
					x:     wm[len(wm)-1].x + 1,
					y:     wm[len(wm)-1].y,
					steps: ws,
				})
			}
		}
	}
	return wm
}

func findIntersectionPoints(wcs [][]coord) []coord {
	var matches []coord
	if len(wcs) == 2 {
		for i := 0; i < len(wcs[0]); i++ {
			for j := 0; j < len(wcs[1]); j++ {
				if wcs[0][i].x == wcs[1][j].x && wcs[0][i].y == wcs[1][j].y {
					matches = append(matches, coord{
						x:     wcs[0][i].x,
						y:     wcs[0][i].y,
						steps: wcs[0][i].steps + wcs[1][j].steps,
					})
				}
			}
		}
	}

	return matches
}

func manhattanDistance(p1, p2 coord) int {
	return int(math.Abs(float64(p1.x-p2.x))) + int(math.Abs(float64(p1.y-p2.y)))
}
