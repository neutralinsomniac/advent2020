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
	modifiedProgram := game.Program{}

	program.InitStateFromFile("input")

	patchIndex := 0
	for {
		modifiedProgram.InitStateFromProgram(&program)
		visitedInstructions := make(map[int]bool)

		// find the first nop or jmp and patch it
		for ; patchIndex < modifiedProgram.Len(); patchIndex++ {
			opcode := modifiedProgram.GetOpcode()
			switch opcode {
			case game.Jmp:
				modifiedProgram.PatchOpcode(patchIndex, game.Nop)
				patchIndex++
				goto Run
			case game.Nop:
				modifiedProgram.PatchOpcode(patchIndex, game.Acc)
				patchIndex++
				goto Run
			}
		}
	Run:
		for {
			ip := modifiedProgram.GetIp()
			if visitedInstructions[ip] {
				break
			}
			visitedInstructions[modifiedProgram.GetIp()] = true
			modifiedProgram.Step()
		}

		if modifiedProgram.GetHalted() {
			fmt.Println(modifiedProgram.GetAcc())
			fmt.Println("time:", time.Since(start))
			return
		}
	}
}
