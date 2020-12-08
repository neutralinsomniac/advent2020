package main

import (
	"bufio"
	"fmt"
	"github.com/golang-collections/go-datastructures/queue"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type bag struct {
	color    string
	count    int
	contains []bag
}

var bags = make(map[string]bag)

func initStateFromFile(filename string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		f := strings.Fields(s)
		var outerbag bag
		outerbag.color = strings.Join(f[0:2], " ")
		if !strings.Contains(s, "no other bags") {
			for i := 4; i < len(f); i += 4 {
				var count int
				var color1, color2 string
				var bag bag
				fmt.Sscanf(strings.Join(f[i:i+4], " "), "%d %s %s bag", &count, &color1, &color2)

				bag.count = count
				bag.color = fmt.Sprintf("%s %s", color1, color2)
				outerbag.contains = append(outerbag.contains, bag)
			}
		}
		bags[outerbag.color] = outerbag
	}
	return
}

func countbags(outerbag bag) int {
	q := queue.New(0)
	q.Put(outerbag)

	sum := 0
	for !q.Empty() {
		b, _ := q.Get(1)
		switch v := b[0].(type) {
		case bag:
			for _, bag := range bags[v.color].contains {
				sum += bag.count
				for i := 0; i < bag.count; i++ {
					q.Put(bags[bag.color])
				}
			}
		}
	}

	return sum
}

func main() {
	start := time.Now()
	initStateFromFile("input")

	sum := countbags(bag{color: "shiny gold"})
	fmt.Println(sum)
	fmt.Println("time:", time.Since(start))
}
