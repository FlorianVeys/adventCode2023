package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	seedKeyValuePair := []int{2906422699,6916147,3075226163,146720986,689152391,244427042,279234546,382175449,1105311711,2036236,3650753915,127044950,3994686181,93904335,1450749684,123906789,2044765513,620379445,1609835129,60050954}
	input, ok := FetchInput("day5-1.txt")
	if !ok {
		panic(errors.New("Unable to fetch input file"))
	}
	game, err := NewGame(input)
	if err != nil {
		panic(err)
	}

	lowestLocation := int(^uint(0) >> 1)
	ProcessSeeds(seedKeyValuePair, func(seed int) {
		result := game.GetCorrespondingAlmanachValue(seed)
		if result < lowestLocation {
			lowestLocation = result
		}
	})

	println(lowestLocation)
}

type AlmanachEntry struct {
	sourceStart int
	destinationStart int
	entryRange int
}
func (entry *AlmanachEntry) GetDestination(source int) (result int, ok bool) {
	if source < entry.sourceStart || source >= entry.sourceStart + entry.entryRange {
		return 0, false
	}

	return entry.destinationStart + (source - entry.sourceStart), true
}
func GetDestination(source int, entries []AlmanachEntry) (result int) {
	for i := 0; i < len(entries); i++ {
		found, ok := entries[i].GetDestination(source)

		if ok == true {
			return found
		}
	}

	return source
}

type Game struct {
	almanach map[string][]AlmanachEntry
	path map[string]string
	startPath string
}
func NewGame(input string) (*Game, error) {
	game := Game{}
	// Hardcoded
	game.startPath = "seed-to-soil"
	game.almanach = make(map[string][]AlmanachEntry)
	game.path = make(map[string]string)

	tokens := strings.Split(input, "\n")

	rName := regexp.MustCompile("([a-z-]+).+:")
	rEntry := regexp.MustCompile("([0-9]+) ([0-9]+) ([0-9]+)")

	currentName := ""

	for _, token := range tokens {
		token = strings.TrimSpace(token)

		if len(token) == 0 {
			continue
		}

		if rName.MatchString(token) {
			name := rName.FindStringSubmatch(token)[1]
			game.almanach[name] = make([]AlmanachEntry, 0)
			if currentName == "" {
				game.path[name] = ""
			} else if name != currentName {
				game.path[currentName] = name
				game.path[name] = ""
			}

			currentName = name
		}

		if rEntry.MatchString(token) {
			matches := rEntry.FindStringSubmatch(token)
			destinationStart, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, err
			}
			sourceStart, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}
			entryRange, err := strconv.Atoi(matches[3])
			if err != nil {
				return nil, err
			}

			game.almanach[currentName] = append(game.almanach[currentName], AlmanachEntry{
				sourceStart: sourceStart,
				destinationStart: destinationStart,
				entryRange: entryRange,
			})
		}
	}

	return &game, nil
}
func (game *Game) GetCorrespondingAlmanachValue(source int) int {
	value := source
	path := game.startPath

	for path != "" {
		value = GetDestination(value, game.almanach[path])
		newPath, ok := game.path[path]

		if ok == true {
			path = newPath
		} else {
			path = ""
		}
	}
	return value
}

// PURE POWER ! NO BRAINS
func ProcessSeeds(keyValuePairs []int, handler func (int)) []int {
	seeds := make([]int, 0)
	keyValue := make([]int, 0)

	for _, value := range keyValuePairs {
		keyValue = append(keyValue, value)

		if len(keyValue) == 2 {
			for i := keyValue[0]; i < keyValue[0] + keyValue[1]; i++ {
				handler(i)
			}

			keyValue = make([]int, 0)
		}

	}

	return seeds
}