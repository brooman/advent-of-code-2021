package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	direction string
	amount    int
}

type position struct {
	horizontal int
	depth      int
	aim        int
}

func getInput() []command {
	raw, _ := ioutil.ReadFile("input")
	data := strings.Split(string(raw), "\n")

	commands := []command{}
	for _, v := range data {
		if len(v) > 0 {
			item := strings.Split(v, " ")
			direction := item[0]
			amount, _ := strconv.Atoi(item[1])

			commands = append(commands, command{
				direction: direction,
				amount:    amount,
			})
		}
	}

	return commands
}

func main() {
	commands := getInput()

	position1 := position{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}

	for _, c := range commands {
		switch c.direction {
		case "forward":
			position1.horizontal += c.amount
			break
		case "down":
			position1.depth += c.amount
			break
		case "up":
			position1.depth -= c.amount
			break
		}

	}

	fmt.Printf("Part 1 answer: %d \n", position1.horizontal*position1.depth)

	// Reset position
	position2 := position{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}

	for _, c := range commands {
		switch c.direction {
		case "forward":
			position2.horizontal += c.amount
			position2.depth += c.amount * position2.aim
			break
		case "down":
			position2.aim += c.amount
			break
		case "up":
			position2.aim -= c.amount
			break
		}
	}

	fmt.Printf("Part 2 answer: %d \n", position2.horizontal*position2.depth)
}
