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

func (c Coord) MultCoord(v Coord) Coord {
	return Coord{row: c.row * v.row, column: c.column * v.column}
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
	pos      Coord
	heading  Direction
	waypoint Coord
}

func (c Coord) Rotate(degrees int) Coord {
	switch degrees {
	case 90:
		return Coord{c.column, -c.row}
	case 180:
		return Coord{-c.row, -c.column}
	case 270:
		return Coord{-c.column, c.row}
	}
	panic("???")

}

func main() {
	start := time.Now()
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	var ship Ship
	ship.heading = East
	ship.waypoint = Coord{row: -1, column: 10}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var direction rune
		var amount int
		fmt.Sscanf(scanner.Text(), "%c%d\n", &direction, &amount)
		fmt.Println("ship:", ship.pos)
		fmt.Println("waypoint:", ship.waypoint)
		fmt.Printf("%c%d\n", direction, amount)
		switch direction {
		case 'N':
			ship.waypoint = ship.waypoint.Add(DirectionToHeading[North].Mult(amount))
		case 'S':
			ship.waypoint = ship.waypoint.Add(DirectionToHeading[South].Mult(amount))
		case 'E':
			ship.waypoint = ship.waypoint.Add(DirectionToHeading[East].Mult(amount))
		case 'W':
			ship.waypoint = ship.waypoint.Add(DirectionToHeading[West].Mult(amount))
		case 'F':
			ship.pos = ship.pos.Add(ship.waypoint.Mult(amount))
		case 'L':
			ship.waypoint = ship.waypoint.Rotate(360 - amount)
		case 'R':
			ship.waypoint = ship.waypoint.Rotate(amount)
		}
	}

	fmt.Println(math.Abs(float64(ship.pos.row)) + math.Abs(float64(ship.pos.column)))
	fmt.Println("time:", time.Since(start))
}
