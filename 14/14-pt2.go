package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Memory map[string]int

func GetAllAddresses(mask string, address string) []string {
	addrs := make([]string, 0)

	//count the number of X's
	var floating []int // indices of X's
	addressSlice := []byte(address)
	for i, b := range mask {
		switch b {
		case 'X':
			floating = append(floating, i)
		case '0':
			addressSlice[i] = address[i]
		case '1':
			addressSlice[i] = '1'
		}
	}
	maxBinaryNum := int(math.Exp2(float64(len(floating))))
	for num := 0; num < maxBinaryNum; num++ {
		numBinary := strconv.FormatInt(int64(num), 2)
		fmtStr := fmt.Sprintf("%%0%ds", len(floating))
		numBinary = fmt.Sprintf(fmtStr, numBinary)
		for i := range numBinary {
			addressSlice[floating[i]] = numBinary[i]
		}
		addrs = append(addrs, string(addressSlice))
	}

	return addrs
}

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
			var addressStr string
			n, _ := fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			if n != 2 {
				panic("couldn't read assignment")
			}

			// zero-pad the address
			addressStr = strconv.FormatInt(int64(address), 2)
			addressStr = fmt.Sprintf("%036s", addressStr)

			addrs := GetAllAddresses(maskStr, addressStr)
			for _, address := range addrs {
				mem[address] = value
			}
		}
	}

	var sum int
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
