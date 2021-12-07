package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(name string) ([]int, error) {
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

	var result []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		number, err := strconv.Atoi(text)
		if err != nil {
			return nil, fmt.Errorf("read input: %w", err)
		}

		result = append(result, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return result, nil
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			result++
		}
	}

	fmt.Println(result) //nolint:forbidigo // 1121
}
