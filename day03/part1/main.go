package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var result []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		result = append(result, text)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return result, nil
}

func binaryStringToInt(input string) int {
	result := 0

	for i := range input {
		result *= 2
		if input[i] == '1' {
			result++
		}
	}

	return result
}

func calcGammaRate(input []string) int {
	zeros := make([]int, len(input[0]))

	for _, line := range input {
		for i := range line {
			if line[i] == '0' {
				zeros[i]++
			}
		}
	}

	result := ""

	for _, count := range zeros {
		if count > len(input)/2 {
			result += "0"
		} else {
			result += "1"
		}
	}

	return binaryStringToInt(result)
}

func calcEpsilonRate(input []string) int {
	zeros := make([]int, len(input[0]))

	for _, line := range input {
		for i := range line {
			if line[i] == '0' {
				zeros[i]++
			}
		}
	}

	result := ""

	for _, count := range zeros {
		if count > len(input)/2 {
			result += "1"
		} else {
			result += "0"
		}
	}

	return binaryStringToInt(result)
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gammaRate := calcGammaRate(input)
	epsilonRate := calcEpsilonRate(input)

	fmt.Println(gammaRate * epsilonRate) //nolint:forbidigo // 3813416
}
