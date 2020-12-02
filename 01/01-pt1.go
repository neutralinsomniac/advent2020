package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
			if arr[i]+arr[j] == 2020 {
				fmt.Println(arr[i] * arr[j])
			}
		}
	}
}
