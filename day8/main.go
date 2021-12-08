package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	patterns []string
	output   []string
}

func getInput() []Entry {
	raw, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")

	entries := []Entry{}
	for _, row := range data {
		if len(row) > 0 {
			delimiter := strings.Split(row, " | ")
			patterns := strings.Split(delimiter[0], " ")
			output := strings.Split(delimiter[1], " ")

			for k, v := range patterns {
				split := strings.Split(v, "")
				sort.Strings(split)
				patterns[k] = strings.Join(split, "")
			}

			for k, v := range output {
				split := strings.Split(v, "")
				sort.Strings(split)
				output[k] = strings.Join(split, "")
			}

			entry := Entry{
				patterns,
				output,
			}

			entries = append(entries, entry)
		}
	}

	return entries
}

func part1() {
	entries := getInput()
	digits := make(map[int]int)

	for _, entry := range entries {
		for _, pattern := range entry.output {
			length := len(pattern)
			switch length {
			case 2:
				digits[1] += 1
			case 3:
				digits[7] += 1
			case 4:
				digits[4] += 1
			case 7:
				digits[8] += 1
			default:
				break
			}
		}
	}

	total := 0
	for _, v := range digits {
		total = total + v
	}

	fmt.Printf("Part 1 answer: %v \n", total)
}

func part2() {
	entries := getInput()

	total := 0
	for _, entry := range entries {
		patterns := append(entry.patterns, entry.output...)

		one := find(patterns, func(item string) bool {
			return len(item) == 2
		})

		four := find(patterns, func(item string) bool {
			return len(item) == 4
		})

		seven := find(patterns, func(item string) bool {
			return len(item) == 3
		})

		eight := find(patterns, func(item string) bool {
			return len(item) == 7
		})

		nine := find(patterns, func(item string) bool {
			return len(item) == 6 && len(intersect(item, four)) == 4
		})

		zero := find(patterns, func(item string) bool {
			return len(item) == 6 &&
				item != nine &&
				len(intersect(item, seven)) == 3
		})

		six := find(patterns, func(item string) bool {
			return len(item) == 6 &&
				item != nine &&
				item != zero
		})

		five := find(patterns, func(item string) bool {
			return len(item) == 5 &&
				len(intersect(item, six)) == 5
		})

		three := find(patterns, func(item string) bool {
			return len(item) == 5 &&
				item != five &&
				len(intersect(item, four)) == 3
		})

		two := find(patterns, func(item string) bool {
			return len(item) == 5 &&
				item != five &&
				item != three
		})

		value := ""
		for _, v := range entry.output {
			switch v {
			case one:
				value += "1"
			case two:
				value += "2"
			case three:
				value += "3"
			case four:
				value += "4"
			case five:
				value += "5"
			case six:
				value += "6"
			case seven:
				value += "7"
			case eight:
				value += "8"
			case nine:
				value += "9"
			case zero:
				value += "0"
			}
		}

		number, _ := strconv.Atoi(value)
		total += number
	}

	fmt.Printf("Part 2 answer: %v \n", total)
}

func find(arr []string, cb func(a string) bool) string {
	for _, v := range arr {
		if cb(v) {
			return v
		}
	}

	return ""
}

func intersect(s1 string, s2 string) []string {
	x := make(map[string]int)
	for _, v := range strings.Split(s1, "") {
		x[v] += 1
	}

	for _, v := range strings.Split(s2, "") {
		x[v] += 1
	}

	res := []string{}
	for k, v := range x {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	part1()
	part2()
}
