package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/5/input")

	numStacks := (len(input[0])-3)/4 + 1
	stacks := make([][]byte, numStacks)

	for i := 0; i < len(input); i++ {
		if len(input[i]) == 0 { // found empty line
			// parse stacks
			for j := i - 2; j >= 0; j-- {
				for k := 0; k < numStacks; k++ {
					spotIndex := (k * 4) + 1
					if input[j][spotIndex] != ' ' { // a crate exists at that spot
						stacks[k] = append(stacks[k], input[j][spotIndex])
					}
				}
			}

			// parse moves
			for i++; i < len(input); i++ {
				re := regexp.MustCompile("[0-9]+")
				moves := re.FindAllString(input[i], -1)
				fromStack, _ := strconv.Atoi(moves[1])
				toStack, _ := strconv.Atoi(moves[2])
				fromStack--
				toStack--
				numToMove, _ := strconv.Atoi(moves[0])
				slice := stacks[fromStack][len(stacks[fromStack])-numToMove:]
				stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-numToMove]
				stacks[toStack] = append(stacks[toStack], slice...)
			}
			break
		}
	}

	result := make([]byte, numStacks)

	for i := 0; i < numStacks; i++ {
		result = append(result, stacks[i][len(stacks[i])-1])
	}

	fmt.Println(string(result[:]))

}
