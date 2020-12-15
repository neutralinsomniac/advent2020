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

type Memory map[int][]int

func (m Memory) Add(index int, value int) {
	if arr, ok := m[index]; ok {
		arr = append(arr, value)
		if len(arr) > 2 {
			arr = arr[1:]
		}
		m[index] = arr
	} else {
		m[index] = make([]int, 1)
		m[index][0] = value
	}
}

func main() {
	start := time.Now()

	file, err := os.Open("input")
	check(err)
	defer file.Close()

	mem := make(Memory, 1000)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	// init
	lastNum := 0
	nums := strings.Split(line, ",")
	for i, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		check(err)
		mem.Add(num, i)
		lastNum = num
	}

	for i := len(nums); i < 30000000; i++ {
		if posArr, ok := mem[lastNum]; ok {
			if len(posArr) == 2 {
				lastNum = posArr[1] - posArr[0]
			} else {
				lastNum = 0
			}
		} else {
			lastNum = 0
		}
		mem.Add(lastNum, i)
		if i == 2020-1 || i == 30000000-1 {
			fmt.Printf("%d number spoken: %d\n", i+1, lastNum)
		}
	}
	fmt.Println("time:", time.Since(start))
}
