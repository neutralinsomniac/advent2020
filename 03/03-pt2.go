package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y int
}

// lol globals
var universeWidth int
var universeHeight int

type Universe map[Coord]bool

func initStateFromFile(filename string) Universe {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	universe := make(Universe)
	var x, y int
	for _, c := range dat {
		switch c {
		case '#':
			universe[Coord{x: x, y: y}] = true
		case '\n':
			y++
			universeWidth = x
			x = 0
			continue
		}
		x++
	}
	universeHeight = y
	return universe
}
func main() {
	universe := initStateFromFile("input")

	mult := 1
	var slopes []Coord
	slopes = append(slopes, Coord{1, 1}, Coord{3, 1}, Coord{5,1}, Coord{7,1}, Coord{1, 2})

	for _, slope := range slopes {
		x := 0
		y := 0
		numTrees := 0

		for y <= universeHeight {
			if universe[Coord{x % universeWidth, y}] {
				numTrees++
			}
			x += slope.x
			y += slope.y
		}
		mult *= numTrees
	}
	fmt.Println(mult)
}
