package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// mul(11,8) <- COrrect syntaxussy

func main() {
	res, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	patters := regex.FindAllStringSubmatch(string(res), -1)
	var nums [][]string

	for _, pattern := range patters {
		regex = regexp.MustCompile("\\d{1,3}")
		temp := regex.FindAllString(pattern[0], -1)

		nums = append(nums, temp)
	}

	total := 0
	for _, n := range nums {
		first, err := strconv.Atoi(n[0])
		second, err := strconv.Atoi(n[1])
		if err != nil {
			panic(err)
		}

		total += first * second
	}

	fmt.Println(total)
}
