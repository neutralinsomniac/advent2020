package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStateFromFile(filename string) (int, []int) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliest, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	idstrings := strings.Split(scanner.Text(), ",")

	ids := make([]int, 0, len(idstrings))
	for _, idstring := range idstrings {
		id, err := strconv.Atoi(idstring)
		if err == nil {
			ids = append(ids, id)
		}
	}

	return earliest, ids
}

func main() {
	start := time.Now()
	earliest, ids := initStateFromFile("input")

	lowestId := 0
	lowestWait := 99999999999
	for _, id := range ids {
		wait := id - (earliest % id)
		if wait < lowestWait {
			lowestId = id
			lowestWait = wait
		}
	}
	fmt.Println(lowestId * lowestWait)
	fmt.Println("time:", time.Since(start))
}
