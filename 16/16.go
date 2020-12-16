package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Restriction struct {
	low1, high1 int
	low2, high2 int
}

type Restrictions map[string]Restriction
type YourTicket []int
type NearbyTickets [][]int

func initStateFromFile(filename string) (Restrictions, YourTicket, NearbyTickets) {

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := 0
	restrictions := make(Restrictions)
	yourticket := make(YourTicket, 0)
	nearbytickets := make(NearbyTickets, 0)
	for {
		scanner.Scan()
		line := scanner.Text()
		switch state {
		case 0:
			if line == "" {
				state = 1
				scanner.Scan() // remove the "your ticket:" line
				continue
			}
			a := strings.Split(line, ":")
			name := a[0]
			var restriction Restriction
			fmt.Sscanf(a[1], "%d-%d or %d-%d", &restriction.low1, &restriction.high1, &restriction.low2, &restriction.high2)
			restrictions[name] = restriction
		case 1:
			if line == "" {
				state = 2
				scanner.Scan() // remove the "nearby tickets:" line
				continue
			}
			nums := strings.Split(line, ",")
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				yourticket = append(yourticket, num)
			}
		case 2:
			if line == "" {
				state = 3
				continue
			}
			nums := strings.Split(line, ",")
			vals := make([]int, 0)
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				vals = append(vals, num)
			}
			nearbytickets = append(nearbytickets, vals)
		default:
			goto End
		}
	}

End:
	return restrictions, yourticket, nearbytickets
}

func main() {
	start := time.Now()

	restrictions, yourticket, nearbytickets := initStateFromFile("input")

	sum := 0
	validtickets := make([][]int, 0, 100)
	validtickets = append(validtickets, yourticket)

NextTicket:
	for _, ticket := range nearbytickets {
	NextValue:
		for _, val := range ticket {
			for _, restriction := range restrictions {
				if (val >= restriction.low1 && val <= restriction.high1) || (val >= restriction.low2 && val <= restriction.high2) {
					continue NextValue
				}
			}
			sum += val
			continue NextTicket
		}

		validtickets = append(validtickets, ticket)
	}

	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))

	start = time.Now()
	// zip dem tickets!
	zipped := make([][]int, 0, len(yourticket))
	for pos := range yourticket {
		arr := make([]int, 0, len(validtickets))
		for _, ticket := range validtickets {
			arr = append(arr, ticket[pos])
		}
		zipped = append(zipped, arr)
	}

	restrictionMap := make(map[string]int, len(restrictions)) // map restriction name to index
	for len(restrictions) != 0 {
		for pos, vals := range zipped {
			numValidRestrictions := 0
			var currentName string
		NextRestriction:
			for name, restriction := range restrictions {
				for _, val := range vals {
					if (val >= restriction.low1 && val <= restriction.high1) || (val >= restriction.low2 && val <= restriction.high2) {
						continue
					}
					continue NextRestriction
				}
				currentName = name
				numValidRestrictions++
			}
			if numValidRestrictions == 1 {
				restrictionMap[currentName] = pos
				delete(restrictions, currentName)
			}
		}
	}

	product := 1
	for name, pos := range restrictionMap {
		if strings.HasPrefix(name, "departure") {
			product *= yourticket[pos]
		}
	}

	fmt.Println(product)

	fmt.Println("time:", time.Since(start))
}
