package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ScoringFunc represents the signature for a scoring function
type ScoringFunc func([]int, []int) int

// readLinesFromFile reads lines from a file and returns them as a slice of strings
func readLinesFromFile(filename string) ([]string, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(input)), "\n"), nil
}

// calculatePoints calculates points for each line using a given scoring function
func calculatePoints(lines []string, scoringFunc ScoringFunc) []int {
	points := make([]int, len(lines))

	for i, line := range lines {
		win, got := parseLine(line)
		points[i] = scoringFunc(win, got)
	}

	return points
}

// calculatePartTwo calculates card counts for part two of the puzzle
func calculatePartTwo(lines []string) []int {
	cardCount := make([]int, len(lines))

	for i := range cardCount {
		cardCount[i] = 1
	}

	for i, line := range lines {
		win, got := parseLine(line)
		points := countIntersection(win, got)

		cardCount = updateCardCount(cardCount, i, points)
	}

	return cardCount
}

// parseLine parses a line into two slices of numbers
func parseLine(s string) ([]int, []int) {
	startAt := strings.Index(strings.TrimSpace(s), ":")
	sets := strings.Split(s[startAt+1:], "|")

	return parseNumbers(sets[0]), parseNumbers(sets[1])
}

// parseNumbers extracts numbers from a string and returns them as a slice of integers
func parseNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	numbers := []int{}

	for _, val := range re.FindAllStringSubmatch(s, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}

// countIntersection counts the intersection of two slices of integers
func countIntersection(set1 []int, set2 []int) int {
	set2Map := make(map[int]bool, len(set2))

	for _, val := range set2 {
		set2Map[val] = true
	}

	intersection := 0
	for _, val := range set1 {
		if set2Map[val] {
			intersection++
		}
	}

	return intersection
}

// calculateTotalPoints calculates the total points from a slice of cards
func calculateTotalPoints(cards []int) int {
	sum := 0

	for _, val := range cards {
		if val > 0 {
			sum += 1 << uint(val-1)
		}
	}

	return sum
}

// calculateTotalCards calculates the total count of cards
func calculateTotalCards(cardCount []int) int {
	sum := 0

	for _, val := range cardCount {
		sum += val
	}

	return sum
}

// updateCardCount updates the count of cards based on given parameters
func updateCardCount(cardCount []int, i, points int) []int {
	updatedCardCount := make([]int, len(cardCount))
	copy(updatedCardCount, cardCount)

	for k := 0; k < cardCount[i]; k++ {
		for j := 1; j <= points; j++ {
			if i+j < len(cardCount) {
				updatedCardCount[i+j]++
			}
		}
	}

	return updatedCardCount
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	lines, err := readLinesFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Part 1
	pointsPart1 := calculatePoints(lines, countIntersection)
	fmt.Println("Part 1 =", calculateTotalPoints(pointsPart1))

	// Part 2
	cardsPart2 := calculatePartTwo(lines)
	fmt.Println("Part 2 =", calculateTotalCards(cardsPart2))
}
