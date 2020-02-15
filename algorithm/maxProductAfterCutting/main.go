package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(solution1(8))
	fmt.Println(solution2(8))
}

func solution1(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}

	products := make([]int, length+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	max := 0
	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
			products[i] = max
		}
	}

	max = products[length]
	return max
}

func solution2(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}

	timesOf3 := length / 3
	if length-timesOf3*3 == 1 {
		timesOf3 -= 1
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}
