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

type Memory map[int]int64

func main() {
	start := time.Now()
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	mem := make(Memory, 1000)
	scanner := bufio.NewScanner(file)

	// each line is either a mask or an assignment
	var maskStr string
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := fmt.Sscanf(line, "mask = %s", &maskStr)
		if n != 1 {
			var address, value int
			n, _ := fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			if n != 2 {
				panic("couldn't read assignment")
			}
			valueStr := strconv.FormatInt(int64(value), 2)
			// zero-pad the value
			valueStr = fmt.Sprintf("%036s", valueStr)
			var newValueStr [36]byte
			for i := range valueStr {
				switch maskStr[i] {
				case 'X':
					newValueStr[i] = valueStr[i]
				default:
					newValueStr[i] = maskStr[i]
				}
			}
			mem[address], err = strconv.ParseInt(string(newValueStr[:]), 2, 64)
			check(err)
		}
	}

	var sum int64
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
