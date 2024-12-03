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

// For part 1
func distanceBetween(sorted1 []int, sorted2 []int) {
	// calculate sum
	var sum int
	for i := 0; i < len(sorted1); i++ {
		diff := math.Abs(float64(sorted1[i] - sorted2[i]))
		sum += int(diff)
	}

	fmt.Println("Total sum:", sum)
}

func intArrayToStringArray(arr []int) []string {
	var output []string

	for _, number := range arr {
		temp := strconv.Itoa(number)
		output = append(output, temp)
	}

	return output
}

// For part 2
func simularityScore(sorted1 []int, sorted2 []int) {
	temp1 := intArrayToStringArray(sorted1)
	temp2 := intArrayToStringArray(sorted2)
	total := 0

	for _, num1 := range temp1 {
		count := 0
		for _, num2 := range temp2 {
			if num1 == num2 {
				count++
			}
		}
		r, e := strconv.Atoi(num1)
		if e != nil {
			panic(e)
		}
		total = total + (r * count)
	}
	fmt.Println("Simularity total: ", total)
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

	// Distance between
	distanceBetween(sorted1, sorted2)

	// Simularity score
	simularityScore(sorted1, sorted2)
}
