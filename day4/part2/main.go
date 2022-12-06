package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input, _ := UrlToLines("https://adventofcode.com/2022/day/4/input")

	re := regexp.MustCompile("[0-9]+")

	count := 0

	for i := 0; i < len(input); i++ {
		reFind := re.FindAllString(input[i], -1)
		var rows []int
		for _, r := range reFind {
			row, _ := strconv.Atoi(r)
			rows = append(rows, row)
		}

		if !(rows[2] > rows[1] || rows[0] > rows[3]) {
			count++
		}
	}

	fmt.Println(count)
}
