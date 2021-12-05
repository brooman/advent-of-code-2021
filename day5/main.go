package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type line struct {
	start point
	end   point
}

type point struct {
	x int
	y int
}

func getInput() []line {
	raw, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(raw), "\n")
	lines := []line{}
	for _, row := range data {
		if len(row) > 0 {
			positions := strings.Split(row, " -> ")
			first := strings.Split(positions[0], ",")
			second := strings.Split(positions[1], ",")

			x1, _ := strconv.Atoi(first[0])
			y1, _ := strconv.Atoi(first[1])

			x2, _ := strconv.Atoi(second[0])
			y2, _ := strconv.Atoi(second[1])

			start := point{
				x: x1,
				y: y1,
			}

			end := point{
				x: x2,
				y: y2,
			}

			line := line{
				start: start,
				end:   end,
			}

			lines = append(lines, line)
		}
	}

	return lines
}

func isIntersecting(p1 point, p2 point, q1 point, q2 point) bool {
	return (((q1.x-p1.x)*(p2.y-p1.y)-(q1.y-p1.y)*(p2.x-p1.x))*((q2.x-p1.x)*(p2.y-p1.y)-(q2.y-p1.y)*(p2.x-p1.x)) < 0) &&
		(((p1.x-q1.x)*(q2.y-q1.y)-(p1.y-q1.y)*(q2.x-q1.x))*((p2.x-q1.x)*(q2.y-q1.y)-(p2.y-q1.y)*(q2.x-q1.x)) < 0)
}

func findIntersection(a1 point, a2 point, b1 point, b2 point) point {
	return point{
		x: ((a1.x*a2.y-a1.y*a2.x)*(b1.x-b2.x) - (a1.x-a2.x)*(b1.x*b2.y-b1.y*b2.x)) / ((a1.x-a2.x)*(b1.y-b2.y) - (a1.y-a2.y)*(b1.x-b2.x)),
		y: ((a1.x*a2.y-a1.y*a2.x)*(b1.y-b2.y) - (a1.y-a2.y)*(b1.x*b2.y-b1.y*b2.x)) / ((a1.x-a2.x)*(b1.y-b2.y) - (a1.y-a2.y)*(b1.x-b2.x)),
	}
}

func isHorizontalOVvertical(line line) bool {
	return line.start.x == line.end.x || line.start.y == line.end.y
}

func part1(lines []line) {
	grid := [][]int{}
	for i := 0; i < 1000; i++ {
		row := []int{}
		for j := 0; j < 1000; j++ {
			row = append(row, 0)
		}

		grid = append(grid, row)
	}

	for _, line := range lines {
		if !isHorizontalOVvertical(line) {
			continue
		}

		var x1 int
		var x2 int
		var y1 int
		var y2 int

		if line.start.x < line.end.x {
			x1 = line.start.x
			x2 = line.end.x
		} else {
			x1 = line.end.x
			x2 = line.start.x
		}

		if line.start.y < line.end.y {
			y1 = line.start.y
			y2 = line.end.y
		} else {
			y1 = line.end.y
			y2 = line.start.y
		}

		for x := x1; x < x2+1; x++ {
			for y := y1; y < y2+1; y++ {
				grid[x][y]++
			}
		}
	}

	atleast2Intersecting := 0
	for _, row := range grid {
		for _, col := range row {
			if col > 1 {
				atleast2Intersecting++
			}
		}
	}

	fmt.Printf("Part 1 answer: %d \n", atleast2Intersecting)
}

func part2(lines []line) {
	grid := [][]int{}
	for i := 0; i < 1000; i++ {
		row := []int{}
		for j := 0; j < 1000; j++ {
			row = append(row, 0)
		}

		grid = append(grid, row)
	}

	for _, line := range lines {
		var xc int
		var yc int

		if line.start.x == line.end.x {
			xc = 0
		} else if line.start.x < line.end.x {
			xc = 1
		} else {
			xc = -1
		}

		if line.start.y == line.end.y {
			yc = 0
		} else if line.start.y < line.end.y {
			yc = 1
		} else {
			yc = -1
		}

		x := line.start.x
		y := line.start.y
		for true {
			grid[x][y]++

			if x == line.end.x && y == line.end.y {
				break
			}

			x = x + xc
			y = y + yc
		}
	}

	atleast2Intersecting := 0
	for _, row := range grid {
		for _, col := range row {
			if col > 1 {
				atleast2Intersecting++
			}
		}
	}

	fmt.Printf("Part 2 answer: %d \n", atleast2Intersecting)
}

func main() {
	lines := getInput()

	part1(lines)
	part2(lines)
}
