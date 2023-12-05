package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
    x  int
    y int
}

func (c Coordinate) toString() string {
	return fmt.Sprintf("%d;%d", c.x, c.y)
}


func main() {
	content, err := os.ReadFile("day3-1.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(content)

	part1(input)

	content2, err := os.ReadFile("day3-2.txt")
	if err != nil {
		fmt.Println(err)
	}
	input2 := string(content2)
	part2(input2)
}

func part1(input string) {
	lines := strings.Split(input, "\n")

	hitMap := getAllNumbersAroundDiscriminant(lines, func(r rune) bool {
		return r != 46 && r != 13 && !isRuneANumber(r)
	})
	numbers := []int{}
	for _, v := range hitMap {
		numbers = append(numbers, reduceCoordinates(v, lines)...)
	}

	result := 0
	for _, v := range numbers {
		result += v
	}
	fmt.Println(result)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	result := 0

	numbersHit := getAllNumbersAroundDiscriminant(lines, func(r rune) bool {
		return r == 42
	})
	
	for _, hits := range numbersHit {
		numbers := reduceCoordinates(hits, lines)

		if len(numbers) == 2 {
			result += numbers[0] * numbers[1]
		}
	}
	println(result)
}

func getAllNumbersAroundDiscriminant(lines []string, discriminant func (r rune) bool) map[string][]Coordinate {
	hitMap := make(map[string][]Coordinate)
	maxX := len(lines[0])
	maxY := len(lines)

	for i, line := range lines {
		for j, char := range line {
			if char != 46 && char != 13 && !isRuneANumber(char) {
				array:= []Coordinate {}
				// Left and right
				appendCoordinateIfANumber(lines, Coordinate{i, j-1}, &array, maxX, maxY)
				appendCoordinateIfANumber(lines, Coordinate{i, j+1}, &array, maxX, maxY)

				// Up and down
				appendCoordinateIfANumber(lines, Coordinate{i-1, j}, &array, maxX, maxY)
				appendCoordinateIfANumber(lines, Coordinate{i+1, j}, &array, maxX, maxY)

				// Diagonals
				appendCoordinateIfANumber(lines, Coordinate{i-1, j-1}, &array, maxX, maxY)
				appendCoordinateIfANumber(lines, Coordinate{i-1, j+1}, &array, maxX, maxY)
				appendCoordinateIfANumber(lines, Coordinate{i+1, j-1}, &array, maxX, maxY)
				appendCoordinateIfANumber(lines, Coordinate{i+1, j+1}, &array, maxX, maxY)
				hitMap[Coordinate{i, j}.toString()] = array
			}
		}
	}
	return hitMap
}

func reduceCoordinates(hitMap []Coordinate, lines []string) []int {
	XYExplored := make(map[string]bool)
	numbers := []int{}

	for _, XY := range hitMap {
		if !mapContainKey(XYExplored, XY.toString()) {
			fullNumber := []string{string(lines[XY.x][XY.y])}
			XYExplored[XY.toString()] = true

			if XY.y + 1 < len(lines[XY.x]) {
				exploreRecursively(Explorer{ 
					line: lines[XY.x],
					y: XY.y + 1,
					maxY: len(lines[XY.x]),
					bounce: 1,
					discover: func(value string, y int) {
						if !mapContainKey(XYExplored, Coordinate{x: XY.x, y: y}.toString()) {
							XYExplored[Coordinate{x: XY.x, y: y}.toString()] = true
							fullNumber = append(fullNumber, value)
						}
					},
				})
			}
			if XY.y - 1 >= 0 {
				exploreRecursively(Explorer{
					line: lines[XY.x],
					y: XY.y - 1,
					maxY: len(lines[XY.x]),
					bounce: -1,
					discover: func(value string, y int) {
						if !mapContainKey(XYExplored, Coordinate{x: XY.x, y: y}.toString()) {
							XYExplored[Coordinate{x: XY.x, y: y}.toString()] = true
							fullNumber = append([]string{value}, fullNumber...)
						}
					},
				})
			}

			number, err := strconv.Atoi(strings.Join(fullNumber, ""))
			if err == nil {
				numbers = append(numbers, number)
			}
		}
	}

	return numbers
}

func isRuneANumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isByteANumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func appendCoordinateIfANumber(lines []string, c Coordinate, hitMap *[]Coordinate, maxX int, maxY int) {
	if c.x < 0 || c.y < 0 || c.x >= maxX || c.y >= maxY {
		return
	}

	if isByteANumber(lines[c.x][c.y]) {
		*hitMap = append(*hitMap, c)
	}
}

type Explorer struct {
	line string
	y int
	maxY int
	bounce int
	discover func(value string, y int)
}
func exploreRecursively(explorer Explorer) {
	if (isByteANumber(explorer.line[explorer.y])) {
		explorer.discover(string(explorer.line[explorer.y]), explorer.y)
		if explorer.y == 0 || explorer.y == explorer.maxY {
			return
		}
		exploreRecursively(Explorer{
			line: explorer.line,
			y: explorer.y + explorer.bounce,
			maxY: explorer.maxY,
			bounce: explorer.bounce,
			discover: explorer.discover,
		})
	}
}

func mapContainKey(mapToCheck map[string]bool, key string) bool {
	_, ok := mapToCheck[key]
	return ok
}