package main

import (
	"fmt"
	"io/ioutil"
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

	days := 80
	ages := input

	for i := 0; i < days; i++ {
		newAges := ages
		for k, v := range ages {
			newValue := v - 1
			if newValue == -1 {
				newAges = append(newAges, 8)
				newValue = 6
			}

			newAges[k] = newValue
		}

		ages = newAges
	}

	fmt.Printf("Part 1 answer: %d \n", len(ages))
}

func part2() {
	input := getInput()

	population := make(map[int]int)
	for _, v := range input {
		population[v] += 1
	}

	days := 256
	for i := 0; i < days; i++ {
		newPopulation := make(map[int]int)
		newPopulation[0] = population[1]
		newPopulation[1] = population[2]
		newPopulation[2] = population[3]
		newPopulation[3] = population[4]
		newPopulation[4] = population[5]
		newPopulation[5] = population[6]
		newPopulation[6] = population[7] + population[0]
		newPopulation[7] = population[8]
		newPopulation[8] = population[0]
		population = newPopulation
	}

	total := 0
	for _, v := range population {
		total = total + v
	}

	fmt.Printf("Part 2 answer: %v \n", total)
}

func main() {
	part1()
	part2()
}
