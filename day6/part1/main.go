package main

import (
	"fmt"
)

func main() {
	lines, _ := UrlToLines("https://adventofcode.com/2022/day/6/input")

	numProcessed := 4
	input := lines[0]
	chars := []byte(input[:4])

	for i := 4; i < len(input); i++ {
		var checkDuplicates map[byte]int
		checkDuplicates = make(map[byte]int)
		foundDuplicate := false

		for _, c := range chars {
			if checkDuplicates[c] != 1 {
				checkDuplicates[c] = 1
			} else {
				foundDuplicate = true
			}
		}

		if !foundDuplicate {
			fmt.Println(string(chars[:]))
			break
		}

		numProcessed++
		chars = chars[1:]
		chars = append(chars, input[i])
	}

	fmt.Println(numProcessed)

}
