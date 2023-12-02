package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
) 

func main1() {
    content, err := os.ReadFile("day1.txt")
    if err != nil {
        fmt.Println(err)
    }
    input := string(content)

	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		result += getFirstLastDigit(line)
	}

	fmt.Println(result)
}

func getFirstLastDigit(s string) int {
	// Find first number in string 
	re := regexp.MustCompile("[0-9]{1}")
	result := re.FindAllString(s, -1)
	
	number, err := strconv.Atoi(result[0] + result[len(result)-1])
	if err != nil {
		fmt.Println(err)
	}
	return number
}