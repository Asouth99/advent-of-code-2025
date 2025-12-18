package day06

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day06/input.txt"
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

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1

	operations := []string{}
	numbersArray := [][]int{}

	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		// logger.Printf("%d : %s", i, line)
		tokens := re.Split(line, -1)
		// logger.Printf("%v", tokens)
		tokensClean := []int{}
		isOperations := false
		for _, token := range tokens {
			// logger.Printf("Token: <%s>", token)
			if token == "" {
				continue
			}
			switch token {
			case "*", "+":
				isOperations = true
				operations = append(operations, token)
			default:
				num, err := strconv.Atoi(token)
				if err != nil {
					logger.Fatalf("error converting string %s to integer: %v\n", token, err)
				}
				tokensClean = append(tokensClean, num)
			}
		}
		if !isOperations {
			numbersArray = append(numbersArray, tokensClean)
		}
	}

	// Print for debugging
	for _, numbers := range numbersArray {
		logger.Printf("Numbers: %v", numbers)
	}
	logger.Printf("Operations: %v", operations)

	// Loop through each column and apply the operation
	for i := 0; i < len(numbersArray[0]); i++ {
		numbersToOperate := []int{}
		operation := operations[i]
		result := 0

		for j := 0; j < len(numbersArray); j++ {
			numbersToOperate = append(numbersToOperate, numbersArray[j][i])
		}

		for _, number := range numbersToOperate {
			switch operation {
			case "*":
				if result == 0 {
					result++
				}
				result *= number
			case "+":
				result += number
			default:
				logger.Fatalf("undefined operation found: %s", operation)
			}
		}
		logger.Printf("Operation %s on %v resulted in %d", operation, numbersToOperate, result)
		answer += result
	}

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++

		// line := scanner.Text()
		// logger.Printf("%d : %s", i, line)
	}
	return answer
}
