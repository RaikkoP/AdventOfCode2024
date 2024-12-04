package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertStringArrayToInt(arr []string) [][]int {
	var result [][]int

	for _, s := range arr {
		var temp []int
		nums := strings.Fields(s)
		for _, n := range nums {
			r, e := strconv.Atoi(n)
			if e != nil {
				panic(e)
			}
			temp = append(temp, r)
		}
		result = append(result, temp)
	}
	return result
}

func checkIfSafe(list []int) bool {
	if len(list) < 2 {
		return true
	}

	// Check increasing sequence
	isIncreasing := true
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		if diff < 1 || diff > 3 {
			isIncreasing = false
			break
		}
	}
	if isIncreasing {
		return true
	}

	// Check decreasing sequence
	isDecreasing := true
	for i := 1; i < len(list); i++ {
		diff := list[i-1] - list[i]
		if diff < 1 || diff > 3 {
			isDecreasing = false
			break
		}
	}
	return isDecreasing
}

func checkIfSafeWithRemoval(list []int) bool {
	if checkIfSafe(list) {
		return true
	}

	for i := 0; i < len(list); i++ {
		modifiedList := make([]int, 0, len(list)-1)
		modifiedList = append(modifiedList, list[:i]...)
		modifiedList = append(modifiedList, list[i+1:]...)

		if checkIfSafe(modifiedList) {
			return true
		}
	}
	return false
}

func main() {
	res, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	arr := convertStringArrayToInt(strings.Split(string(res), "\n"))
	count := 0

	for _, list := range arr {
		if checkIfSafeWithRemoval(list) {
			count++
		}
	}
	fmt.Println(count)
}
