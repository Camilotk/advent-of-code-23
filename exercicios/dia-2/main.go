package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CubeCounter struct to manage cube counts and related operations
type CubeCounter struct {
	MaxCounts map[string]int // Maximum counts of cubes for each color
}

// NewCubeCounter initializes a CubeCounter object
func NewCubeCounter() *CubeCounter {
	return &CubeCounter{
		MaxCounts: map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		},
	}
}

// CalculatePower calculates the power based on maximum counts of cubes for each color
func (cc *CubeCounter) CalculatePower() int {
	result := 1
	for _, value := range cc.MaxCounts {
		result *= value
	}
	return result
}

// UpdateMaxCount updates the maximum count of cubes for a specific color if the provided number is greater
func (cc *CubeCounter) UpdateMaxCount(number int, color string) {
	if number > cc.MaxCounts[color] {
		cc.MaxCounts[color] = number
	}
}

// IsCountPossible checks if the provided count for a color is within the limit
func (cc *CubeCounter) IsCountPossible(count int, color string) bool {
	switch color {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	default:
		return false
	}
}

// ProcessGameData processes each line of game data to update counts and determine if a game is valid
func (cc *CubeCounter) ProcessGameData(line string) int {
	isValidGame := true
	gameID := strings.Split(line, ":")[0]
	gameID = strings.Split(gameID, "Game ")[1]

	entries := strings.Split(line, ":")[1]
	games := strings.Split(entries, ";")
	for _, game := range games {
		rounds := strings.Split(game, ",")
		for _, round := range rounds {
			data := strings.TrimSpace(round)
			set := strings.Split(data, " ")
			count, _ := strconv.Atoi(set[0])
			color := set[1]
			isPossible := cc.IsCountPossible(count, color)
			if !isPossible {
				isValidGame = false
			}
			cc.UpdateMaxCount(count, color)
		}
	}

	if isValidGame {
		id, _ := strconv.Atoi(gameID)
		return id
	}
	return 0
}

func main() {
	totalValidGames := 0
	totalPower := 0
	cc := NewCubeCounter()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalValidGames += cc.ProcessGameData(scanner.Text())
		totalPower += cc.CalculatePower()

		// Reset max counts for the next game
		cc.MaxCounts["red"] = 0
		cc.MaxCounts["green"] = 0
		cc.MaxCounts["blue"] = 0
	}

	// Part 1
	fmt.Println("Total Valid Games:", totalValidGames)
	// Part 2
	fmt.Println("Total Power:", totalPower)
}
