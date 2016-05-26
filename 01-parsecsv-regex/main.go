package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var parseRegexp *regexp.Regexp = regexp.MustCompile(`\s*(,|$)\s*`)

func main() {
	list, err := parseCSV("1,2,3,4,5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", list)
}

func parseCSV(s string) ([]int, error) {
	var list []int
	matches := parseRegexp.Split(s, -1)
	for _, match := range matches {
		if match == "" {
			continue
		}
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		list = append(list, num)
	}
	return list, nil
}
