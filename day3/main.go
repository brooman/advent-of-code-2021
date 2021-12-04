package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() []string {
	raw, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	binaryArray := []string{}
	for _, binary := range data {
		if len(binary) > 0 {
			binaryArray = append(binaryArray, binary)
		}
	}

	return binaryArray
}

func main() {
	binaryArray := getInput()

	values := [12]int{}
	for _, row := range binaryArray {
		for k, v := range strings.Split(row, "") {
			value, _ := strconv.Atoi(v)
			values[k] = values[k] + value
		}
	}

	var mostCommon string
	var leastCommon string

	for _, v := range values {
		if v > (len(binaryArray) / 2) {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}

	}

	gammaRate, _ := strconv.ParseInt(mostCommon, 2, 64)
	epsilonRate, _ := strconv.ParseInt(leastCommon, 2, 64)
	powerConsumption := gammaRate * epsilonRate

	fmt.Printf("Part 1 answer: %d \n", powerConsumption)

	oxygenRatings := binaryArray
	co2Ratings := binaryArray

	currentOxygenIndex := 0
	for len(oxygenRatings) > 1 {
		newOxygenRatings := []string{}

		zeros := 0
		ones := 0
		for _, v := range oxygenRatings {
			value, _ := strconv.Atoi(string(v[currentOxygenIndex]))
			if value == 0 {
				zeros++
			} else {
				ones++
			}
		}

		var mostCommon int
		if ones >= zeros {
			mostCommon = 1
		} else {
			mostCommon = 0
		}

		for _, v := range oxygenRatings {
			if string(v[currentOxygenIndex]) == strconv.Itoa(mostCommon) {
				newOxygenRatings = append(newOxygenRatings, v)
			}
		}

		if len(newOxygenRatings) > 1 {
			currentOxygenIndex++
		}

		oxygenRatings = newOxygenRatings
	}

	currentCo2Index := 0
	for len(co2Ratings) > 1 {
		newCo2Ratings := []string{}

		zeros := 0
		ones := 0
		for _, v := range co2Ratings {
			value, _ := strconv.Atoi(string(v[currentCo2Index]))
			if value == 0 {
				zeros++
			} else {
				ones++
			}
		}

		var leastCommon int
		if ones >= zeros {
			leastCommon = 0
		} else {
			leastCommon = 1
		}

		for _, v := range co2Ratings {
			if string(v[currentCo2Index]) == strconv.Itoa(leastCommon) {
				newCo2Ratings = append(newCo2Ratings, v)
			}
		}

		if len(newCo2Ratings) > 1 {
			currentCo2Index++
		}

		co2Ratings = newCo2Ratings
	}

	oxygenRating, _ := strconv.ParseInt(oxygenRatings[0], 2, 64)
	co2Rating, _ := strconv.ParseInt(co2Ratings[0], 2, 64)
	lifeSupportRating := oxygenRating * co2Rating

	fmt.Printf("Part 2 answer: %d \n", lifeSupportRating)
}
