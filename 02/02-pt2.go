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
		var char byte
		var password string
		fmt.Sscanf(s, "%d-%d %c: %s", &low, &high, &char, &password)
		var i int
		if password[low-1] == char {
			i++
		}
		if password[high-1] == char {
			i++
		}
		if i == 1 {
			numValidPasswords++
		}
	}

	fmt.Println(numValidPasswords)
}
