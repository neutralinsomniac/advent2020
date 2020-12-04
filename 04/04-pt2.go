package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Passport map[string]string

var hclre = regexp.MustCompile("^#[0-9a-f]{6}$")
var eclre = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var pidre = regexp.MustCompile("^[0-9]{9}$")

func (p Passport) isValid() bool {
	byr, err := strconv.Atoi(p["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr, err := strconv.Atoi(p["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, err := strconv.Atoi(p["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	var height int
	var unit string
	fmt.Sscanf(p["hgt"], "%d%s", &height, &unit)
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}

	if !hclre.MatchString(p["hcl"]) {
		return false
	}

	if !eclre.MatchString(p["ecl"]) {
		return false
	}

	if !pidre.MatchString(p["pid"]) {
		return false
	}
	return true
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
			spl := strings.Split(field, ":")
			passport[spl[0]] = spl[1]
		}
		passports = append(passports, passport)
	}
	return passports
}

func main() {
	passports := initStateFromFile("input")

	numValid := 0
	for _, passport := range passports {
		if passport.isValid() {
			numValid++
		}
	}
	fmt.Println(numValid)
}
