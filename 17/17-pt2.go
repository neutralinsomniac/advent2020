package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y, z, w int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{x: c.x + other.x, y: c.y + other.y, z: c.z + other.z, w: c.w + other.w}
}

func (c Coord) Mult(v int) Coord {
	return Coord{x: c.x * v, y: c.y * v, z: c.z * v, w: c.w * v}
}

type Cell int

const (
	Dead Cell = iota
	Alive
)

func (s Cell) String() string {
	str, _ := CellToChar[s]
	return string(str)
}

var CharToCell = map[rune]Cell{'.': Dead, '#': Alive}
var CellToChar = map[Cell]rune{Dead: '.', Alive: '#'}

type Grid map[Coord]Cell

func initStateFromFile(filename string) Grid {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	grid := make(Grid)
	scanner := bufio.NewScanner(file)
	x := 0
	z := 0
	w := 0
	for scanner.Scan() {
		for y, cell := range scanner.Text() {
			grid[Coord{x: x, y: y, z: z, w: w}] = CharToCell[cell]
		}
		x++
	}

	return grid
}

func generateAdjacentCoords() {
	adjacentCoords = make([]Coord, 0, 80)

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						adjacentCoords = append(adjacentCoords, Coord{x, y, z, w})
					}
				}
			}
		}
	}
}

var adjacentCoords []Coord

func countAliveAdjacent(grid Grid, coord Coord) int {
	alive := 0

	for _, adj := range adjacentCoords {
		if (grid)[coord.Add(adj)] == Alive {
			alive++
		}
	}

	return alive
}

func copyGrid(old Grid) Grid {
	newGrid := make(Grid, len(old))
	for k, v := range old {
		newGrid[k] = v
	}
	return newGrid
}

func calcNewState(grid Grid) (Grid, bool) {
	newGrid := make(Grid, len(grid))
	var changed bool

	// expand cells to consider
	expandedGrid := copyGrid(grid)
	for coord := range expandedGrid {
		for _, adj := range adjacentCoords {
			if _, ok := grid[coord.Add(adj)]; !ok {
				expandedGrid[coord.Add(adj)] = Dead
			}
		}
	}

	for coord, cell := range expandedGrid {
		numAlive := countAliveAdjacent(expandedGrid, coord)
		switch cell {
		case Alive:
			if numAlive == 2 || numAlive == 3 {
				newGrid[coord] = Alive
			} else {
				newGrid[coord] = Dead
			}
		case Dead:
			if numAlive == 3 {
				newGrid[coord] = Alive
			}
		}
	}
	return newGrid, changed
}

func countAlive(grid Grid) int {
	alive := 0
	for _, cell := range grid {
		if cell == Alive {
			alive++
		}
	}
	return alive
}

func main() {
	start := time.Now()
	init := initStateFromFile("input")

	generateAdjacentCoords()
	grid := init

	for i := 0; i < 6; i++ {
		grid, _ = calcNewState(grid)
	}

	fmt.Println(countAlive(grid))
	fmt.Println("time:", time.Since(start))
}
