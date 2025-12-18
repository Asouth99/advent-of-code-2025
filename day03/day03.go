package day03

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day03/input.txt"
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
	for scanner.Scan() {
		i++
		batteryBank := scanner.Text()

		// logger.Printf("%d : %s", i, batteryBank)

		// Find largest digit in bank - that will be our first number.
		// Find largest digit in the bank to the right of our first digit - that will be our second number
		maxN1 := 0
		maxN1Index := 0
		for i := 0; i < len(batteryBank)-1; i++ {
			digit, err := strconv.Atoi(string(batteryBank[i]))
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", string(batteryBank[i]), err)
			}
			if digit > maxN1 {
				maxN1 = digit
				maxN1Index = i
			}
		}

		maxN2 := 0
		// maxN2Index := 0
		for i := maxN1Index + 1; i < len(batteryBank); i++ {
			digit, err := strconv.Atoi(string(batteryBank[i]))
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", string(batteryBank[i]), err)
			}
			if digit > maxN2 {
				maxN2 = digit
				// maxN2Index = i
			}
		}

		maxJoltage := 10*maxN1 + maxN2
		answer += maxJoltage
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
		batteryBank := scanner.Text()
		logger.Printf("%d : %s", i, batteryBank)

		// Joltage is now 12 digits instead of 2
		joltage := 0
		maxIndex := -1
		for i := 1; i <= 12; i++ { // Loop for every digit of the 12
			maxDigit := 0
			logger.Printf("Checking %s for digit %d in our joltage", batteryBank[maxIndex+1:len(batteryBank)-(12-i)], i)
			for j := maxIndex + 1; j < len(batteryBank)-(12-i); j++ { // Loop through the rest of the batteryBank
				digit, err := strconv.Atoi(string(batteryBank[j]))
				if err != nil {
					logger.Fatalf("error converting string %s to integer: %v\n", string(batteryBank[i]), err)
				}
				if digit > maxDigit {
					maxDigit = digit
					maxIndex = j
				}
			}
			joltage += int(math.Pow10(12-i)) * maxDigit
			logger.Printf("MaxDigit: %d, MaxIndex: %d, Joltage: %d", maxDigit, maxIndex, joltage)
		}
		logger.Printf("Joltage: %d", joltage)
		answer += joltage
	}
	return answer
}
