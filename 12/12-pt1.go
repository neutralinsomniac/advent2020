package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	row, column int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{row: c.row + other.row, column: c.column + other.column}
}

func (c Coord) Mult(v int) Coord {
	return Coord{row: c.row * v, column: c.column * v}
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	}
	return "??"
}

var DirectionToHeading = map[Direction]Coord{
	North: Coord{row: -1, column: 0},
	East:  Coord{row: 0, column: 1},
	South: Coord{row: 1, column: 0},
	West:  Coord{row: 0, column: -1},
}

type Ship struct {
	pos     Coord
	heading Direction
}

func main() {
	start := time.Now()
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	var ship Ship
	ship.heading = East

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction rune
		var amount int
		fmt.Sscanf(scanner.Text(), "%c%d\n", &direction, &amount)
		switch direction {
		case 'N':
			ship.pos = ship.pos.Add(DirectionToHeading[North].Mult(amount))
		case 'S':
			ship.pos = ship.pos.Add(DirectionToHeading[South].Mult(amount))
		case 'E':
			ship.pos = ship.pos.Add(DirectionToHeading[East].Mult(amount))
		case 'W':
			ship.pos = ship.pos.Add(DirectionToHeading[West].Mult(amount))
		case 'F':
			ship.pos = ship.pos.Add(DirectionToHeading[ship.heading].Mult(amount))
		case 'L':
			ship.heading = Direction((int(ship.heading) + int((360-amount)/90)) % 4)
		case 'R':
			ship.heading = Direction((int(ship.heading) + int(amount/90)) % 4)
		}
	}

	fmt.Println(math.Abs(float64(ship.pos.row)) + math.Abs(float64(ship.pos.column)))
	fmt.Println("time:", time.Since(start))
}
