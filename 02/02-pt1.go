package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numValidPasswords int
	for scanner.Scan() {
		s := scanner.Text()
		var low, high int
		var char int32
		var password string
		fmt.Sscanf(s, "%d-%d %c: %s", &low, &high, &char, &password)
		var i int
		for _, c := range password {
			if c == char {
				i++
			}
		}
		if i >= low && i <= high {
			numValidPasswords++
		}
	}

	fmt.Println(numValidPasswords)
}
