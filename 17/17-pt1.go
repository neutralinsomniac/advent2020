package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y, z int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{x: c.x + other.x, y: c.y + other.y, z: c.z + other.z}
}

func (c Coord) Mult(v int) Coord {
	return Coord{x: c.x * v, y: c.y * v, z: c.z * v}
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

func (grid Grid) String() string {
	sb := strings.Builder{}

	maxX := -99999999
	maxY := -99999999
	minX := 99999999
	minY := 99999999
	zs := make(map[int]bool)
	// first find max x, y and all z's
	for coord := range grid {
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
		zs[coord.z] = true
	}

	zKeys := make([]int, 0, len(zs))
	for k := range zs {
		zKeys = append(zKeys, k)
	}
	sort.Ints(zKeys)
	for _, z := range zKeys {
		sb.WriteString(fmt.Sprintf("z=%d\n", z))
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				cell := grid[Coord{x, y, z}]
				sb.WriteRune(CellToChar[cell])
			}
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

func initStateFromFile(filename string) Grid {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	grid := make(Grid)
	scanner := bufio.NewScanner(file)
	x := 0
	z := 0
	for scanner.Scan() {
		for y, cell := range scanner.Text() {
			grid[Coord{x: x, y: y, z: z}] = CharToCell[cell]
		}
		x++
	}

	return grid
}

var adjacentCoords = []Coord{
	Coord{x: 0, y: -1, z: -1},
	Coord{x: 0, y: -1, z: 0},
	Coord{x: 0, y: -1, z: 1},

	Coord{x: 0, y: 0, z: -1},
	Coord{x: 0, y: 0, z: 1},

	Coord{x: 0, y: 1, z: -1},
	Coord{x: 0, y: 1, z: 0},
	Coord{x: 0, y: 1, z: 1},

	Coord{x: 1, y: -1, z: -1},
	Coord{x: 1, y: -1, z: 0},
	Coord{x: 1, y: -1, z: 1},

	Coord{x: 1, y: 0, z: -1},
	Coord{x: 1, y: 0, z: 0},
	Coord{x: 1, y: 0, z: 1},

	Coord{x: 1, y: 1, z: -1},
	Coord{x: 1, y: 1, z: 0},
	Coord{x: 1, y: 1, z: 1},

	Coord{x: -1, y: -1, z: -1},
	Coord{x: -1, y: -1, z: 0},
	Coord{x: -1, y: -1, z: 1},

	Coord{x: -1, y: 0, z: -1},
	Coord{x: -1, y: 0, z: 0},
	Coord{x: -1, y: 0, z: 1},

	Coord{x: -1, y: 1, z: -1},
	Coord{x: -1, y: 1, z: 0},
	Coord{x: -1, y: 1, z: 1},
}

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
	for coord, cell := range grid {
		if cell == Alive {
			for _, adj := range adjacentCoords {
				if _, ok := grid[coord.Add(adj)]; !ok {
					grid[coord.Add(adj)] = Dead
				}
			}
		}
	}

	for coord, cell := range grid {
		numAlive := countAliveAdjacent(grid, coord)
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

	grid := init

	for i := 0; i < 6; i++ {
		grid, _ = calcNewState(grid)
	}

	fmt.Println(countAlive(grid))
	fmt.Println("time:", time.Since(start))
}
