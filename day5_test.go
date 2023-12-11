package main

import (
	"reflect"
	"testing"
)

func FixtureNewGameExample() (*Game, error) {
	input:= `seed-to-soil map:
	50 98 2
	52 50 48
	
	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15
	
	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4
	
	water-to-light map:
	88 18 7
	18 25 70
	
	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13
	
	temperature-to-humidity map:
	0 69 1
	1 0 69
	
	humidity-to-location map:
	60 56 37
	56 93 4`

	return NewGame(input)
}

func TestNewGame(t *testing.T) {
	game, err := FixtureNewGameExample()

	if err != nil {
		t.Error(err.Error())
	}

	expectSeedToSoil := []AlmanachEntry{ AlmanachEntry{destinationStart: 50, sourceStart: 98, entryRange: 2}, AlmanachEntry{destinationStart: 52, sourceStart: 50, entryRange: 48}}
	if !reflect.DeepEqual(game.almanach["seed-to-soil"], expectSeedToSoil) {
		t.Errorf("Expected %v, got %v", expectSeedToSoil, game.almanach["seed-to-soil"])
	}

	expectSoilToFertilizer := []AlmanachEntry{ AlmanachEntry{destinationStart: 0, sourceStart: 15, entryRange: 37}, AlmanachEntry{destinationStart: 37, sourceStart: 52, entryRange: 2}, AlmanachEntry{destinationStart: 39, sourceStart: 0, entryRange: 15}}
	if !reflect.DeepEqual(game.almanach["soil-to-fertilizer"], expectSoilToFertilizer) {
		t.Errorf("Expected %v, got %v", expectSoilToFertilizer, game.almanach["soil-to-fertilizer"])
	}

	expectFertilizerToWater := []AlmanachEntry{ AlmanachEntry{destinationStart: 49, sourceStart: 53, entryRange: 8}, AlmanachEntry{destinationStart: 0, sourceStart: 11, entryRange: 42}, AlmanachEntry{destinationStart: 42, sourceStart: 0, entryRange: 7}, AlmanachEntry{destinationStart: 57, sourceStart: 7, entryRange: 4} }
	if !reflect.DeepEqual(game.almanach["fertilizer-to-water"], expectFertilizerToWater) {
		t.Errorf("Expected %v, got %v", expectFertilizerToWater, game.almanach["fertilizer-to-water"])
	}

	expectWaterToLight := []AlmanachEntry{ AlmanachEntry{destinationStart: 88, sourceStart: 18, entryRange: 7}, AlmanachEntry{destinationStart: 18, sourceStart: 25, entryRange: 70} }
	if !reflect.DeepEqual(game.almanach["water-to-light"], expectWaterToLight) {
		t.Errorf("Expected %v, got %v", expectWaterToLight, game.almanach["water-to-light"])
	}

	expectLightToTemperature := []AlmanachEntry{ AlmanachEntry{destinationStart: 45, sourceStart: 77, entryRange: 23}, AlmanachEntry{destinationStart: 81, sourceStart: 45, entryRange: 19}, AlmanachEntry{destinationStart: 68, sourceStart: 64, entryRange: 13} }
	if !reflect.DeepEqual(game.almanach["light-to-temperature"], expectLightToTemperature) {
		t.Errorf("Expected %v, got %v", expectLightToTemperature, game.almanach["light-to-temperature"])
	}

	expectTemperatureToHumidity := []AlmanachEntry{ AlmanachEntry{destinationStart: 0, sourceStart: 69, entryRange: 1}, AlmanachEntry{destinationStart: 1, sourceStart: 0, entryRange: 69} }
	if !reflect.DeepEqual(game.almanach["temperature-to-humidity"], expectTemperatureToHumidity) {
		t.Errorf("Expected %v, got %v", expectTemperatureToHumidity, game.almanach["temperature-to-humidity"])
	}

	expectHumidityToLocation := []AlmanachEntry{ AlmanachEntry{destinationStart: 60, sourceStart: 56, entryRange: 37}, AlmanachEntry{destinationStart: 56, sourceStart: 93, entryRange: 4} }
	if !reflect.DeepEqual(game.almanach["humidity-to-location"], expectHumidityToLocation) {
		t.Errorf("Expected %v, got %v", expectHumidityToLocation, game.almanach["humidity-to-location"])
	}

	if game.path["seed-to-soil"] != "soil-to-fertilizer" {
		t.Errorf("Expected %v, got %v", "soil-to-fertilizer", game.path["seed-to-soil"])
	}

	if game.path["soil-to-fertilizer"] != "fertilizer-to-water" {
		t.Errorf("Expected %v, got %v", "fertilizer-to-water", game.path["soil-to-fertilizer"])
	}

	if game.path["fertilizer-to-water"] != "water-to-light" {
		t.Errorf("Expected %v, got %v", "water-to-light", game.path["fertilizer-to-water"])
	}

	if game.path["water-to-light"] != "light-to-temperature" {
		t.Errorf("Expected %v, got %v", "light-to-temperature", game.path["water-to-light"])
	}

	if game.path["light-to-temperature"] != "temperature-to-humidity" {
		t.Errorf("Expected %v, got %v", "temperature-to-humidity", game.path["light-to-temperature"])
	}

	if game.path["temperature-to-humidity"] != "humidity-to-location" {
		t.Errorf("Expected %v, got %v", "humidity-to-location", game.path["temperature-to-humidity"])
	}
}

func TestAlmanachEntryGetDestinationBelowRange(t *testing.T) {
	entry := AlmanachEntry{destinationStart: 50, sourceStart: 98, entryRange: 2}
	
	source := 97

	_, ok := entry.GetDestination(source)

	if ok == true {
		t.Errorf("should not found destination")
	}
}
func TestAlmanachEntryGetDestinationAboveRange(t *testing.T) {
	entry := AlmanachEntry{destinationStart: 50, sourceStart: 98, entryRange: 2}
	
	source := 100

	_, ok := entry.GetDestination(source)

	if ok == true {
		t.Errorf("should not found destination")
	}
}

func TestAlmanachEntryGetDestinationInRange1(t *testing.T) {
	entry := AlmanachEntry{destinationStart: 50, sourceStart: 98, entryRange: 5}
	
	source := []int {98, 99, 100, 101, 102}
	expect := []int {50, 51, 52, 53, 54}

	for i := 0; i < len(source); i++ {
		result, ok := entry.GetDestination(source[i])

		if ok == false {
			t.Errorf("should found destination %v, from source %v", expect[i], source[i])
		}

		if result != expect[i] {
			t.Errorf("Expected %v, got %v", expect[i], result)
		}
	}
}
func TestAlmanachEntryGetDestinationInRange2(t *testing.T) {
	entry := AlmanachEntry{destinationStart: 0, sourceStart: 15, entryRange: 37}
	
	source := []int {15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	expect := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

	for i := 0; i < len(source); i++ {
		result, ok := entry.GetDestination(source[i])

		if ok == false {
			t.Errorf("should found destination %v, from source %v", expect[i], source[i])
		}

		if result != expect[i] {
			t.Errorf("Expected %v, got %v", expect[i], result)
		}
	}
}
func TestAlmanachEntryGetDestinationInRange3(t *testing.T) {
	entry := AlmanachEntry{ destinationStart: 39, sourceStart: 0, entryRange: 15 }
	
	source := 1000

	_, ok := entry.GetDestination(source)

	if ok == true {
		t.Errorf("should not found %v", source)
	}
}

func TestAlmanachEntriesGetDestinationNotInRanges(t *testing.T) {
	entries := []AlmanachEntry{
		{ destinationStart: 0, sourceStart: 15, entryRange: 37 },
		{ destinationStart: 37, sourceStart: 52, entryRange: 2 },
		{ destinationStart: 39, sourceStart: 0, entryRange: 15 },
	}

	source := 1000
	expect := 1000

	result := GetDestination(source, entries)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestAlmanachEntriesGetDestinationInRanges(t *testing.T) {
	entries := []AlmanachEntry{
		{ destinationStart: 0, sourceStart: 15, entryRange: 37 },
		{ destinationStart: 37, sourceStart: 52, entryRange: 2 },
		{ destinationStart: 39, sourceStart: 0, entryRange: 15 },
	}

	source := [] int{0, 10, 52, 54, 15, 20}
	expect := [] int{39, 49, 37, 54, 0, 5}

	for i := 0; i < len(source); i++ {
		result := GetDestination(source[i], entries)

		if result != expect[i] {
			t.Errorf("Expected %v, got %v", expect[i], result)
		}
	}
}

func TestGameGetCorrespondingAlmanachValue(t *testing.T) {
	game, err := FixtureNewGameExample()

	if err != nil {
		t.Error(err.Error())
	}

	source := []int {79, 14, 55, 13}
	destination := []int {82, 43, 86, 35}

	for i := 0; i < len(source); i++ {
		result := game.GetCorrespondingAlmanachValue(source[i])

		if result != destination[i] {
			t.Errorf("Expected %v, got %v", destination[i], result)
		}
	}
}

func TestSeedList(t *testing.T) {
	input := []int{ 0, 3 }
	expect := []int {0, 1, 2}

	result := GenerateSeedsList(input)

	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Errorf("Expected %v, got %v", expect[i], result[i])
		}
	}
}

func TestSeedList2(t *testing.T) {
	input := []int{ 79, 14, 55, 13 }
	expect := []int {79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67}

	result := GenerateSeedsList(input)

	for i := 0; i < len(result); i++ {
		if result[i] != expect[i] {
			t.Errorf("Expected %v, got %v", expect[i], result[i])
		}
	}
}