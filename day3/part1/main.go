package main

import (
	"fmt"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/3/input")

	var bothCompartments []byte

	for i := 0; i < len(input); i++ {
		var compartment1 map[byte]int
		compartment1 = make(map[byte]int)
		for j := 0; j < len(input[i])/2; j++ {
			compartment1[input[i][j]] = 1
		}
		for j := len(input[i]) / 2; j < len(input[i]); j++ {
			if compartment1[input[i][j]] == 1 {
				bothCompartments = append(bothCompartments, input[i][j])
				break
			}
		}
	}
	total := 0
	for i := 0; i < len(bothCompartments); i++ {
		if 'Z' >= bothCompartments[i] {
			total += int((bothCompartments[i] - 'A') + 27)
		} else {
			total += int((bothCompartments[i] - 'a') + 1)
		}
	}
	fmt.Println(total)
}
