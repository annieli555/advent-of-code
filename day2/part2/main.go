package main

import (
	"fmt"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/2/input")
	// file, _ := os.Open("../input.txt")
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// var input []string

	// for scanner.Scan() {
	// 	input = append(input, scanner.Text())
	// }

	totalScore := 0

	for i := 0; i < len(input); i++ {
		// get opponent's move
		opponent := input[i][0] - 'A'

		// get result (0 = lose, 1 = draw, 2 = win)
		result := input[i][2] - 'X'
		score := result * 3

		// result == 0 --> 1 less than opponent
		// result == 1 --> same as opponent
		// result == 2 --> 1 more than opponent
		// calculate my move
		me := (opponent + result - 1 + 3) % 3

		score += me + 1

		totalScore += int(score)
	}
	fmt.Println(totalScore)
}
