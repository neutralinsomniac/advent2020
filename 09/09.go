package main

import (
	"bufio"
	"fmt"
	"os"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	start := time.Now()
	ints := initStateFromFile("input")

	num := 0

NextNumber1:
	for i := 0; i < len(ints)-25; i++ {
		check := ints[i : i+25]
		num = ints[i+25]
		// now figure out if num == check[x]+check[y]
		for x := range check {
			for y := range check[1:] {
				if check[x]+check[y] == num {
					continue NextNumber1
				}
			}
		}
		fmt.Println(num)
		fmt.Println("time:", time.Since(start))
		break
	}

	start = time.Now()

NextNumber2:
	for i := range ints {
		sum := ints[i]
		smallest := ints[i]
		largest := ints[i]
		for _, j := range ints[i+1:] {
			sum += j
			smallest = min(smallest, j)
			largest = max(largest, j)
			if sum == num {
				fmt.Println(largest + smallest)
				fmt.Println("time:", time.Since(start))
				return
			}
			if j > num {
				continue NextNumber2
			}
		}
	}
}
