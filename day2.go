package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
) 

func main() {
	constraints := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
    content, err := os.ReadFile("day2.txt")
    if err != nil {
        fmt.Println(err)
    }
    input := string(content)

	lines := strings.Split(input, "\n")

	gameIdSum := 0

	for _, line := range lines {
		r := regexp.MustCompile(`Game (?P<ID>[\d]+):(?P<rules>.+)`)
		isPossible := true
		
		gameIdMatches := r.FindStringSubmatch(line)
		gameId, err := strconv.Atoi(gameIdMatches[1])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		r2 := regexp.MustCompile(`(?P<amount>[\d]+) (?P<color>[\w]+)`)
		
		r2Matches := r2.FindAllStringSubmatch(line, -1)

		for _, match := range r2Matches {
			number, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			if number > constraints[match[2]] {
				isPossible = false
				break
			}
		}

		if isPossible {
			gameIdSum += gameId
		}
	}

	fmt.Println(gameIdSum)
}