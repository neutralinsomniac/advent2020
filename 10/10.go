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
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	ints := make([]int, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}

	return ints
}

// length 1, 2, 3
var precalcSums = []int{1, 1, 2}

func getPrecalcSum(length int) int {
	if length <= len(precalcSums) {
		return precalcSums[length-1]
	}
	for i := len(precalcSums); i < length; i++ {
		precalcSums = append(precalcSums, precalcSums[i-1]+precalcSums[i-2]+precalcSums[i-3])
	}
	return precalcSums[length-1]
}
func main() {
	start := time.Now()
	ints := initStateFromFile("input")
	// add the outlet
	ints = append(ints, 0)

	sort.Ints(ints)

	// add the device
	ints = append(ints, ints[len(ints)-1]+3)

	threes := 0
	ones := 0
	for i := range ints[:len(ints)-1] {
		if ints[i+1]-ints[i] == 1 {
			ones++
		} else if ints[i+1]-ints[i] == 3 {
			threes++
		}
	}
	fmt.Println(ones * threes)
	fmt.Println("time:", time.Since(start))

	start = time.Now()

	product := 1
	for start := 0; start < len(ints)-1; {
		for end := start; end < len(ints)-1; end++ {
			if ints[end+1]-ints[end] == 3 {
				sum := getPrecalcSum(end - start + 1)
				product *= sum
				start = end + 1
				break
			}
		}
	}

	fmt.Println(product)
	fmt.Println("time:", time.Since(start))
}
