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

	program.InitStateFromFile("input")

	patchIndex := 0
	for {
		visitedInstructions := make(map[int]bool)

		// find the first nop or jmp and patch it
		for ; patchIndex < program.Len(); patchIndex++ {
			opcode := program.GetOpcode(patchIndex)
			switch opcode {
			case game.Jmp:
				program.PatchOpcode(patchIndex, game.Nop)
				goto Run
			case game.Nop:
				program.PatchOpcode(patchIndex, game.Jmp)
				goto Run
			}
		}
	Run:
		for {
			ip := program.GetIp()
			if visitedInstructions[ip] {
				break
			}
			visitedInstructions[program.GetIp()] = true
			program.Step()
		}

		if program.GetHalted() {
			fmt.Println(program.GetAcc())
			fmt.Println("time:", time.Since(start))
			return
		}

		// undo the previous patch
		opcode := program.GetOpcode(patchIndex)
		switch opcode {
		case game.Jmp:
			program.PatchOpcode(patchIndex, game.Nop)
		case game.Nop:
			program.PatchOpcode(patchIndex, game.Jmp)
		}
		program.Reset()
		// inc patchIndex here to avoid repeating the same patch we just performed
		patchIndex++
	}
}
