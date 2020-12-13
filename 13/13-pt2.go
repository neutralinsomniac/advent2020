package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Bus struct {
	q int
	r int
}

type ByQ []Bus

func (a ByQ) Len() int {
	return (len(a))
}

func (a ByQ) Less(i, j int) bool {
	return a[i].q < a[j].q
}

func (a ByQ) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func initStateFromFile(filename string) []Bus {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// only care about the second line
	scanner.Scan()
	idstrings := strings.Split(scanner.Text(), ",")

	buses := make([]Bus, 0)
	for i, idstring := range idstrings {
		id, err := strconv.Atoi(idstring)
		if err == nil {
			buses = append(buses, Bus{q: id, r: i})
		}
	}

	sort.Sort(sort.Reverse(ByQ(buses)))
	return buses
}

func main() {
	start := time.Now()
	buses := initStateFromFile("input")

	multiple := buses[0].q
Next:
	for t := buses[0].q - buses[0].r; ; t += multiple {
		for i, bus := range buses {
			if (t+bus.r)%bus.q != 0 {
				continue Next
			}
			if i != 0 {
				multiple = 1
				for j := 0; j < i+1; j++ {
					multiple *= buses[j].q
				}
			}
		}
		fmt.Println("answer:", t)
		break
	}

	fmt.Println("time:", time.Since(start))
}
