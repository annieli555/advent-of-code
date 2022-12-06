package main

import (
	"fmt"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/3/input")

	var groups []byte

	for i := 0; i < len(input); i++ {
		var group1 map[byte]int
		group1 = make(map[byte]int)
		var group2 map[byte]int
		group2 = make(map[byte]int)
		for j := 0; j < len(input[i]); j++ {
			group1[input[i][j]] = 1
		}
		for j := 0; j < len(input[i+1]); j++ {
			group2[input[i+1][j]] = 1
		}
		i += 2
		for j := 0; j < len(input[i]); j++ {
			if group1[input[i][j]] == 1 && group2[input[i][j]] == 1 {
				groups = append(groups, input[i][j])
				break
			}
		}
	}
	total := 0
	for i := 0; i < len(groups); i++ {
		if 'Z' >= groups[i] {
			total += int((groups[i] - 'A') + 27)
		} else {
			total += int((groups[i] - 'a') + 1)
		}
	}
	fmt.Println(total)
}
