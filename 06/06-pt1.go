package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStateFromFile(filename string) []map[rune]bool {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	var answers []map[rune]bool
	for _, group := range strings.Split(string(dat), "\n\n") {
		answer := make(map[rune]bool)
		for _, c := range group {
			if c != '\n' {
				answer[c] = true
			}
		}
		answers = append(answers, answer)
	}

	return answers
}

func main() {
	start := time.Now()
	groups := initStateFromFile("input")

	sum := 0
	for _, group := range groups {
		sum += len(group)
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
