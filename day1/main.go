package main

import (
	"fmt"
	"strconv"
)

func main() {
	calories, _ := UrlToLines("https://adventofcode.com/2022/day/1/input")
	// file, _ := os.Open("./input.txt")
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// var calories []string

	// for scanner.Scan() {
	// 	calories = append(calories, scanner.Text())
	// }
	calories = append(calories, "")

	currSum := 0
	highest := 0
	secondHighest := 0
	thirdHighest := 0
	for i := 0; i < len(calories); i++ {
		if len(calories[i]) > 0 {
			food, _ := strconv.Atoi(calories[i])
			currSum += food
		} else {
			if highest < currSum {
				thirdHighest = secondHighest
				secondHighest = highest
				highest = currSum
			} else if secondHighest < currSum {
				thirdHighest = secondHighest
				secondHighest = currSum
			} else if thirdHighest < currSum {
				thirdHighest = currSum
			}
			currSum = 0
		}
	}
	fmt.Println(highest, secondHighest, thirdHighest)
	fmt.Println(highest + secondHighest + thirdHighest)
}
