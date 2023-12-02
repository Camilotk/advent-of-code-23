package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Part 1
func getFirstAndLastDigits(line string) (int, int) {
	firstDigit := -1
	lastDigit := -1

	for _, c := range line {
		if unicode.IsDigit(c) {
			digit, _ := strconv.Atoi(string(c))
			if firstDigit == -1 {
				firstDigit = digit
			}
			lastDigit = digit
		}
	}

	return firstDigit, lastDigit
}

func getSumOfValues(filename string) int {
	sum := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return sum
	}
	defer file.Close()

	var lines []string
	var line string
	for {
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	for _, line := range lines {
		firstDigit, lastDigit := getFirstAndLastDigits(line)

		if firstDigit != -1 {
			sum = sum + firstDigit*10 + lastDigit
		}
	}

	return sum
}

// Part 2
var numberWords = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3r",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func extractNumbers(line string) string {
	for word, digit := range numberWords {
		line = strings.ReplaceAll(line, word, digit)
	}
	return line
}

func extractCalibrationNumber(line string) (int, error) {
	var digits []byte
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			digits = append(digits, line[i])
		}
	}
	number, err := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
	return number, err
}

func getSumOfCalibrationValues(filename string) (int, error) {
	sum := 0

	file, err := os.Open(filename)
	if err != nil {
		return sum, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("Line:", line)

		line = extractNumbers(line)
		number, err := extractCalibrationNumber(line)
		if err != nil {
			return sum, err
		}
		// fmt.Println("Calibration Number:", number)

		firstDigit := number / 10
		lastDigit := number % 10

		if firstDigit != 0 {
			sum += firstDigit*10 + lastDigit
		}
	}

	if err := scanner.Err(); err != nil {
		return sum, err
	}

	return sum, nil
}

// main
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Part 1
	sum := getSumOfValues(filename)
	fmt.Println("Sum part 1: ", sum)

	// Part 2
	sum2, err := getSumOfCalibrationValues(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Sum part 2: ", sum2)
}
