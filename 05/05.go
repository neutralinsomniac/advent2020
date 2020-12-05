package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStateFromFile(filename string) []int {
	f, err := os.Open(filename)
	check(err)

	var seats []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		seat := []byte(scanner.Text())
		for i, c := range seat {
			switch c {
			case 'F', 'L':
				seat[i] = '0'
			case 'B', 'R':
				seat[i] = '1'
			}
		}
		i, _ := strconv.ParseInt(string(seat), 2, 32)
		seats = append(seats, int(i))
	}

	return seats
}

func main() {
	start := time.Now()
	seats := initStateFromFile("input")

	sort.Ints(seats)

	fmt.Println("part 1: ", seats[len(seats)-1])

	for i := range seats[:len(seats)-1] {
		if seats[i] != seats[i+1]-1 {
			fmt.Println("part 2: ", seats[i]+1)
			fmt.Println("time:", time.Since(start))
			return
		}
	}
}
