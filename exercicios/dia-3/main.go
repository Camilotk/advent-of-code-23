package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// node represents a position in the schema grid.
type node struct {
	row int
	col int
}

// Schematics holds the schema data and the sums.
//
//	schemas is the input schema grid split by rows
//	schemaSum is the sum of extracted numbers from the schema
//	gearRatioSum is the sum of product of two extracted numbers marked by '*'
type Schematics struct {
	schemas      []string
	schemaSum    int
	gearRatioSum int
}

// NewSchematicsFromFile creates a new Schematics instance from a file.
func NewSchematicsFromFile(filename string) (*Schematics, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %v", err)
	}

	return &Schematics{
		schemas: strings.Split(string(file), "\n"),
	}, nil
}

// processSchemas traverses the schema grid and updates sums based on symbols.
func (s *Schematics) processSchemas() {
	for i := range s.schemas {
		for j := range s.schemas[i] {
			if isSymbol(s.schemas[i][j]) {
				s.updateSchemaSum(node{i, j})
			}
		}
	}
}

// isSymbol checks if a character is not a letter, digit, or period.
func isSymbol(char byte) bool {
	return !unicode.IsLetter(rune(char)) && !unicode.IsDigit(rune(char)) && char != '.'
}

// traverseAndFindNum traverses from a given node and extracts a number from the schema.
func (s *Schematics) traverseAndFindNum(startNode node) int {
	schemaToCheck := s.schemas[startNode.row]
	leftI, rightI := startNode.col, startNode.col

	for i := startNode.col - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(schemaToCheck[i])) {
			break
		}
		leftI = i
	}

	for i := startNode.col + 1; i < len(schemaToCheck); i++ {
		if !unicode.IsDigit(rune(schemaToCheck[i])) {
			break
		}
		rightI = i
	}

	resultNum, err := strconv.Atoi(schemaToCheck[leftI : rightI+1])
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}

	dotStr := strings.Repeat(".", rightI-leftI+1)
	s.schemas[startNode.row] = schemaToCheck[:leftI] + dotStr + schemaToCheck[rightI+1:]

	return resultNum
}

// updateSchemaSum calculates the sum and gear ratio sum based on the marked nodes.
func (s *Schematics) updateSchemaSum(symbolNode node) {
	rDelta := []int{0, -1, -1, -1, 0, 1, 1, 1}
	cDelta := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	var numbers []int

	for i := range rDelta {
		nNode := node{
			row: symbolNode.row + rDelta[i],
			col: symbolNode.col + cDelta[i],
		}

		if isValidNode(nNode, s) && unicode.IsDigit(rune(s.schemas[nNode.row][nNode.col])) {
			numbers = append(numbers, s.traverseAndFindNum(nNode))
		}
	}

	for _, num := range numbers {
		s.schemaSum += num
	}

	if len(numbers) == 2 && rune(s.schemas[symbolNode.row][symbolNode.col]) == '*' {
		s.gearRatioSum += numbers[0] * numbers[1]
	}
}

// isValidNode checks if a node position is within schema bounds.
func isValidNode(n node, s *Schematics) bool {
	return 0 <= n.row && n.row < len(s.schemas) && 0 <= n.col && n.col < len(s.schemas[0])
}

// main
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	s, err := NewSchematicsFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s.processSchemas()

	// Part 1
	fmt.Printf("The final sum is %d \n", s.schemaSum)
	// Part 2
	fmt.Printf("The final gear ratio sum is %d", s.gearRatioSum)
}
