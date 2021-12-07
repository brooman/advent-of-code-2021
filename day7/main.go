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
	for _, v := range strings.Split(data[0], ",") {

		age, _ := strconv.Atoi(v)
		values = append(values, age)
	}

	return values
}

func part1() {
	input := getInput()

	lowest := math.MaxInt32
	highest := 0
	for _, v := range input {
		if v > highest {
			highest = v
		}

		if v < lowest {
			lowest = v
		}
	}

	leastDistance := math.MaxInt32
	for i := lowest; i < highest+1; i++ {
		totalDistance := 0
		for _, v := range input {
			totalDistance += int(math.Abs(float64(v) - float64(i)))
		}

		if totalDistance < leastDistance {
			leastDistance = totalDistance
		}
	}

	fmt.Printf("Part 1 answer: %v \n", leastDistance)
}

func part2() {
	input := getInput()

	lowest := math.MaxInt32
	highest := 0
	for _, v := range input {
		if v > highest {
			highest = v
		}

		if v < lowest {
			lowest = v
		}
	}

	leastDistance := math.MaxInt32
	for i := lowest; i < highest+1; i++ {
		totalDistance := 0
		for _, v := range input {
			distance := int(math.Abs(float64(v) - float64(i)))
			for i := 0; i < distance+1; i++ {
				totalDistance += i
			}
		}

		if totalDistance < leastDistance {
			leastDistance = totalDistance
		}
	}

	fmt.Printf("Part 2 answer: %v \n", leastDistance)
}

func main() {
	part1()
	part2()
}
