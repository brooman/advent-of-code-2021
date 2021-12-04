package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getInput() []int {
	raw, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")
	values := []int{}
	for _, v := range data {
		number, _ := strconv.Atoi(v)
		if number > 0 {
			values = append(values, number)
		}
	}

	return values
}

func main() {
	values := getInput()

	increases := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increases++
		}
	}

	fmt.Printf("Part 1 answer: %d \n", increases)

	sumIncreases := 0
	prevSum := math.MaxInt32
	for i := 2; i < len(values); i++ {
		sum := values[i] + values[i-1] + values[i-2]

		if sum > prevSum {
			sumIncreases++
		}

		prevSum = sum
	}

	fmt.Printf("Part 2 answer: %d \n", sumIncreases)
}
