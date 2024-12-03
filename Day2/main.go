package main

import (
	"fmt"
	"math"
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

func main() {
	res, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	arr := convertStringArrayToInt(strings.Split(string(res), "\n"))
	count := 0

	for _, list := range arr {
		var trend int
		if list[0] > list[1] {
			trend = -1 // Decreasing
		} else if list[0] < list[1] {
			trend = 1 // Increasing
		} else {
			continue
		}
		for i := 2; i < len(list); i++ {
			l := list[i-2] // Left
			m := list[i-1] // Middle
			r := list[i]   // Right

			if trend == 1 {
				if l > m || m > r {
					break
				}
			} else if trend == -1 {
				if l < m || m < r {
					break
				}
			}

			lDif := math.Abs(float64(l - m))
			rDif := math.Abs(float64(m - r))

			if lDif > 3 || lDif < 1 || rDif > 3 || rDif < 1 {
				break
			}

			if i == len(list)-1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
