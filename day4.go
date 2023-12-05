package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("day4-1.txt")
	if err != nil {
		fmt.Println(err)
	}
	input := string(content)


	scratch, err := NewScratchRandomDraw(input)
	if err != nil {
		panic(err)
	}


	fmt.Println(scratch.toString())
	fmt.Println(scratch.sumPoints())
}

type Card struct {
	result string
	cardId int
	winningNumbers []int
	points int
}

func NewScratchRandomDraw(result string) (*ScratchRandomDraw, error) {
	scratchRandomDraw := new(ScratchRandomDraw)
	loadError := scratchRandomDraw.loadResult(result)
	if (loadError != nil) {
		return nil, loadError
	}
	
	searchWinningError := scratchRandomDraw.searchWinningNumbersForEachCard()
	if (searchWinningError != nil) {
		return nil, searchWinningError
	}

	calculatePointsError := scratchRandomDraw.calculatePoints()
	if (calculatePointsError != nil) {
		return nil, calculatePointsError
	}

	return scratchRandomDraw, nil
}
type ScratchRandomDraw struct {
	cards []Card
}
func (s *ScratchRandomDraw) loadResult(result string) error {
	lines := strings.Split(result, "\n")
	s.cards = []Card{}

	for _, cardRow := range lines {
		r := regexp.MustCompile(`Card\s+([\d]+):(.+)`)
		extract := r.FindStringSubmatch(cardRow)

		if (len(extract) != 3) {
			return errors.New("Invalid card row formatting")
		}

		cardId, err := strconv.Atoi(extract[1])
		if (err != nil) {
			return errors.New(fmt.Sprintf("Unable to parse card id %s from %v", extract[1], extract))
		}

		s.cards = append(s.cards, Card{
			result: extract[2],
			cardId: cardId,
			winningNumbers: []int{},
		})
	}
	return nil
}
func (s *ScratchRandomDraw) searchWinningNumbersForEachCard() error {
	for index, card := range s.cards {
		winningResult := strings.Split(card.result, "|")
		if (len(winningResult) != 2) {
			return errors.New("Invalid card row formatting")
		}

		winningResult[0] = strings.TrimSpace(winningResult[0])
		winningResult[1] = strings.TrimSpace(winningResult[1])

		winningNumbers := strings.Split(winningResult[0], " ")
		resultNumbers := strings.Split(winningResult[1], " ")

		for _, resultNumber := range resultNumbers {
			if isStringANumber(resultNumber) && slices.Contains(winningNumbers, resultNumber) {
				numberParsed, err := strconv.Atoi(resultNumber)
				if (err != nil) {
					return errors.New(fmt.Sprintf("Unable to parse number %s", resultNumber))
				}
				s.cards[index].winningNumbers = append(s.cards[index].winningNumbers, numberParsed)
			}
		}
	}
	return nil
}
func (s *ScratchRandomDraw) toString() string {
	output := "[\n"
	for _, card := range s.cards {
		output += "{\n result:" + card.result + "\n cardId:" + strconv.Itoa(card.cardId) + "\n winningNumbers:" + fmt.Sprint(card.winningNumbers) + "\n points: " + strconv.Itoa(card.points) + "\n}\n"
	}
	output += "]"
	return output
}
func (s *ScratchRandomDraw) calculatePoints() error {
	for index, card := range s.cards {
		if len(card.winningNumbers) > 0 {
			binaryString := "1" + strings.Repeat("0", len(card.winningNumbers) - 1)

			if i, err := strconv.ParseInt(binaryString, 2, 64); err == nil {
				s.cards[index].points = int(i)
			} else {
				return errors.New(fmt.Sprint("Unable to calculate point for %s", card.cardId))
			}
		}
	}

	return nil
}
func (s *ScratchRandomDraw) sumPoints() int {
	sum := 0
	for _, card := range s.cards {
		sum += card.points
	}
	return sum
}

func isStringANumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}