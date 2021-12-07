package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type inputLine struct {
	command string
	units   int
}

func readInput(name string) ([]inputLine, error) {
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

	var result []inputLine

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		units, err := strconv.Atoi(text[1])
		if err != nil {
			return nil, fmt.Errorf("read input: %w", err)
		}

		result = append(result, inputLine{command: text[0], units: units})
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

	position := 0
	depth := 0
	aim := 0

	for _, line := range input {
		switch line.command {
		case "forward":
			position += line.units
			depth += aim * line.units
		case "down":
			aim += line.units
		case "up":
			aim -= line.units
		}
	}

	result := position * depth

	fmt.Println(result) //nolint:forbidigo // 1971232560
}
