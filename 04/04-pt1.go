package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Passport map[string]bool

func (p Passport) isValid() bool {
	if p["byr"] && p["iyr"] && p["eyr"] && p["hgt"] && p["hcl"] && p["ecl"] && p["pid"] {
		return true
	}
	return false
}

func initStateFromFile(filename string) []Passport {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	passportsStr := strings.Split(string(dat), "\n\n")
	passports := make([]Passport, 0)

	for _, passportStr := range passportsStr {
		passport := Passport{}
		fields := strings.Fields(passportStr)
		for _, field := range fields {
			key := strings.Split(field, ":")[0]
			passport[key] = true
		}
		passports = append(passports, passport)
	}
	return passports
}

func main() {
	passports := initStateFromFile("input")

	numValid := 0
	for _, passport := range passports {
		fmt.Println(passport)
		fmt.Println("---")
		if passport.isValid() {
			numValid++
		}
	}
	fmt.Println(numValid)
}
