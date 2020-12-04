package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, num)
	}
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					fmt.Println(arr[i] * arr[j] * arr[k])
				}
			}
		}
	}
	fmt.Println("time:", time.Since(start))
}
