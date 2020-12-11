package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Seat int

const (
	Floor Seat = iota
	Empty
	Occupied
)

func (s Seat) String() string {
	str, _ := SeatToChar[s]
	return string(str)
}

var CharToSeat = map[rune]Seat{'.': Floor, '#': Occupied, 'L': Empty}
var SeatToChar = map[Seat]rune{Floor: '.', Occupied: '#', Empty: 'L'}

type Grid map[Coord]Seat

func (grid Grid) String() string {
	row := 0
	column := 0
	sb := strings.Builder{}

	for {
		seat, ok := grid[Coord{row, column}]
		if !ok {
			if column == 0 {
				break
			}
			row++
			column = 0
			sb.WriteString("\n")
			continue
		}
		sb.WriteRune(SeatToChar[seat])
		column++
	}

	return sb.String()
}

func initStateFromFile(filename string) Grid {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	grid := make(Grid)
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		for column, seat := range scanner.Text() {
			grid[Coord{row: row, column: column}] = CharToSeat[seat]
		}
		row++
	}

	return grid
}

var adjacentCoords = []Coord{
	Coord{row: 0, column: 1},
	Coord{row: 0, column: -1},
	Coord{row: 1, column: 1},
	Coord{row: 1, column: -1},
	Coord{row: 1, column: 0},
	Coord{row: -1, column: 0},
	Coord{row: -1, column: 1},
	Coord{row: -1, column: -1},
}

func countOccupiedAdjacent(grid Grid, coord Coord) int {
	occupied := 0

	for _, adj := range adjacentCoords {
		if (grid)[coord.Add(adj)] == Occupied {
			occupied++
		}
	}

	return occupied
}

func countOccupiedLOS(grid Grid, coord Coord) int {
	occupied := 0

NextDirection:
	for _, adj := range adjacentCoords {
		for vector := 1; ; vector++ {
			seat, ok := grid[coord.Add(adj.Mult(vector))]
			if !ok || seat == Empty {
				// reached end of map
				continue NextDirection
			}
			if seat == Occupied {
				occupied++
				continue NextDirection
			}
		}
	}
	return occupied
}

func calcNewState(grid Grid) (Grid, bool) {
	newGrid := make(Grid, len(grid))
	var changed bool

	for coord, seat := range grid {
		adj := countOccupiedAdjacent(grid, coord)
		if seat == Empty && adj == 0 {
			newGrid[coord] = Occupied
			changed = true
		} else if seat == Occupied && adj >= 4 {
			newGrid[coord] = Empty
			changed = true
		} else {
			newGrid[coord] = seat
		}
	}
	return newGrid, changed
}

func calcNewState2(grid Grid) (Grid, bool) {
	newGrid := make(Grid, len(grid))
	var changed bool

	for coord, seat := range grid {
		adj := countOccupiedLOS(grid, coord)
		if seat == Empty && adj == 0 {
			newGrid[coord] = Occupied
			changed = true
		} else if seat == Occupied && adj >= 5 {
			newGrid[coord] = Empty
			changed = true
		} else {
			newGrid[coord] = seat
		}
	}
	return newGrid, changed
}

func countOccupied(grid Grid) int {
	occupied := 0
	for _, seat := range grid {
		if seat == Occupied {
			occupied++
		}
	}
	return occupied
}

func main() {
	start := time.Now()
	init := initStateFromFile("input")

	grid := init

	changed := false
	for {
		grid, changed = calcNewState(grid)
		if !changed {
			break
		}
	}

	fmt.Println(countOccupied(grid))
	fmt.Println("time:", time.Since(start))

	start = time.Now()

	grid = init

	changed = false
	for {
		grid, changed = calcNewState2(grid)
		if !changed {
			break
		}
	}

	fmt.Println(countOccupied(grid))

	fmt.Println("time:", time.Since(start))
}
