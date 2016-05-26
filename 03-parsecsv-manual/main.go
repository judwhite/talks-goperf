package main

import (
	"fmt"
	"log"
)

func main() {
	list, err := parseCSV("1,2,3,4,5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", list)
}

func parseCSV(s string) ([]int, error) {
	list := make([]int, 0, len(s)/2+1)

	inDigit := false
	num := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			inDigit = true
			num = num*10 + int(c-'0')
		} else {
			if inDigit {
				list = append(list, num)
				num = 0
				inDigit = false
			}
		}
	}

	if inDigit {
		list = append(list, num)
	}

	return list, nil
}
