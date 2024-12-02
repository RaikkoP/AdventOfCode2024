package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Simple bubble sort for integers
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func convertFromStringToInt(arr [][]string) []int {
	var output []int
	for _, num := range arr {
		// Join the digits and convert to a single integer
		numStr := strings.Join(num, "")
		n, e := strconv.Atoi(numStr)
		if e != nil {
			panic(e)
		}
		output = append(output, n)
	}
	return output
}

func splitArray(content string) ([]int, []int) {
	lines := strings.Split(content, "\n")
	var temp1 [][]string
	var temp2 [][]string

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			temp1 = append(temp1, strings.Split(numbers[0], ""))
			temp2 = append(temp2, strings.Split(numbers[1], ""))
		}
	}

	return convertFromStringToInt(temp1), convertFromStringToInt(temp2)
}

func main() {
	var array1, array2 []int

	// Read the input file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	array1, array2 = splitArray(string(content))

	sorted1 := bubbleSort(array1)
	sorted2 := bubbleSort(array2)

	// calculate sum
	var sum int
	for i := 0; i < len(sorted1); i++ {
		diff := math.Abs(float64(sorted1[i] - sorted2[i]))
		sum += int(diff)
		fmt.Printf("sorted1[%d]: %d, sorted2[%d]: %d, Difference: %.0f\n",
			i, sorted1[i], i, sorted2[i], diff)
	}

	fmt.Println("Total sum:", sum)
}
