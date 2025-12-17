package day01

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day01/input.txt"
	if len(inputFile) > 0 {
		file = inputFile[0]
	}

	switch part {
	case 1:
		return SolvePart1(file, logger), nil
	case 2:
		return SolvePart2(file, logger), nil
	default:
		return -1, errors.New("incorrect part number recieved")
	}
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dialPos := 50 // Dial starts at position 50
	answer := 0

	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			logger.Fatalf("error converting string %d to integer: %v\n", number, err)
		}
		// logger.Printf("%d: %s%d", i, direction, number)
		dialStart := dialPos // dialStart only used for logging
		switch direction {
		case "L":
			dialPos = ((dialPos-number)%100 + 100) % 100 // modulo in Go gives remainder even for negative numbers
		case "R":
			dialPos = ((dialPos+number)%100 + 100) % 100
		default:
			logger.Fatalf("incorrect direction '%s' recieved", direction)
		}
		logger.Printf("Dial has been rotated by %s%02d from %02d->%02d", direction, number, dialStart, dialPos)

		// Update answer if dial is poiting to 0
		if dialPos == 0 {
			answer++
		}
	}
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	dialPos := 50 // Dial starts at position 50
	answer := 0

	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		direction := string(line[0])
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			logger.Fatalf("error converting string %d to integer: %v\n", number, err)
		}
		// logger.Printf("%d: %s%d", i, direction, number)
		dialStart := dialPos // Take note of the starting dial position
		switch direction {
		case "L":
			dialPos = ((dialPos-number)%100 + 100) % 100
		case "R":
			dialPos = ((dialPos+number)%100 + 100) % 100
		default:
			logger.Fatalf("incorrect direction '%s' recieved", direction)
		}

		// Update answer by however many times we pass 0
		numFullRotations := number / 100
		numZerosPassed := numFullRotations
		remainder := number % 100
		if (direction == "R" && dialStart < 100 && (dialStart+remainder) > 100) || (direction == "L" && dialStart > 0 && (dialStart-remainder) < 0) {
			numZerosPassed += 1
		}
		answer += numZerosPassed

		// Update answer if dial is poiting to 0
		if dialPos == 0 {
			answer++
		}

		// Logging
		logger.Printf("Dial has been rotated by %s%02d from %02d->%02d", direction, number, dialStart, dialPos)
		logger.Printf("Rotations:%d | NumZerosPassed:%d | answer:%d", numFullRotations, numZerosPassed, answer)
	}
	return answer
}
