package day06

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	numbersArray := []string{}
	operationIndexes := []int{}
	operations := ""

	re := regexp.MustCompile(`[\+\*]`)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		switch string(line[0]) {
		case "*", "+":
			operations = line
			for _, val := range re.FindAllStringIndex(line, -1) {
				operationIndexes = append(operationIndexes, val[0])
			}
			logger.Printf("Operations: %s", operations)
			logger.Printf("OperationIndexes: %v", operationIndexes)
		default:
			numbersArray = append(numbersArray, line)
		}
	}

	// Print for debugging
	for _, numbers := range numbersArray {
		logger.Printf("Numbers: %s", numbers)
	}

	// Loop through each operation. This gives the starting and ending index for the numbers
	for i := 0; i < len(operationIndexes); i++ {
		// Get how many digits in the numbers
		opIndex := operationIndexes[i]
		digitLength := 0
		if i == (len(operationIndexes) - 1) {
			digitLength = len(operations) - opIndex
		} else {
			digitLength = operationIndexes[i+1] - opIndex - 1
		}
		logger.Printf("DigitLength: %d", digitLength)

		// Get the numbers that we need to operate on
		numbersToOperate := []int{}
		for j := 0; j < digitLength; j++ {
			index := opIndex + j
			// Create number string
			numberString := ""
			for _, numbers := range numbersArray {
				numberString = numberString + string(numbers[index])
			}
			numberString = strings.ReplaceAll(numberString, " ", "") // Remove all spaces from number
			number, err := strconv.Atoi(numberString)
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", numberString, err)
			}
			// Add number to numbersToOperate
			numbersToOperate = append(numbersToOperate, number)
		}

		// Operate on each number in numbersToOperate
		operation := string(operations[opIndex])
		result := 0
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
				logger.Fatalf("incorrect operation %s found", operation)
			}
		}

		// debugging
		logger.Printf("NumbersToOperate: %v, Op: %s, Result: %d", numbersToOperate, operation, result)
		answer += result

	}

	return answer
}
