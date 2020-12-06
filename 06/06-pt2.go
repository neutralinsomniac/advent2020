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

func initStateFromFile(filename string) []map[rune]int {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	var answers []map[rune]int
	for _, group := range strings.Split(string(dat), "\n\n") {
		answer := make(map[rune]int)
		for _, c := range strings.Trim(group, "\n") {
			answer[c] += 1
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
		for _, count := range group {
			if count == group['\n']+1 {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
