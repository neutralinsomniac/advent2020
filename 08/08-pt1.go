package main

import (
	"fmt"
	"github.com/neutralinsomniac/advent2020/game"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	program := game.Program{}

	visitedInstructions := make(map[int]bool)

	program.InitStateFromFile("input")

	for {
		ip := program.GetIp()
		if visitedInstructions[ip] {
			fmt.Printf("%d\n", program.GetAcc())
			fmt.Println("time:", time.Since(start))
			return
		}
		visitedInstructions[program.GetIp()] = true
		program.Step()
	}
}
