package day10

import (
	internal "aoc2025/internal/utils"
	"bufio"
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day10/input.txt"
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

	indicators := []map[int]int{}
	allButtons := [][][]int{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		// Read indicator
		indicator := map[int]int{}
		for pos, char := range lineSplit[0][1 : len(lineSplit[0])-1] {
			var n int
			switch char {
			case '.':
				n = 0
			case '#':
				n = 1
			default:
				logger.Fatalf("incorrect character <%s> at index %d found in indicator %v", string(char), pos, lineSplit[0][1:len(lineSplit[0])-1])
			}
			indicator[pos] = n
		}
		indicators = append(indicators, indicator)

		// Read buttons
		buttons := [][]int{}
		for _, button := range lineSplit[1 : len(lineSplit)-1] {
			b := []int{}
			for _, char := range button[1 : len(button)-1] {
				if char == ',' {
					continue
				}
				num, err := strconv.Atoi(string(char))
				if err != nil {
					logger.Fatalf("error converting string %s to integer: %v\n", string(char), err)
				}
				b = append(b, num)
			}
			buttons = append(buttons, b)
		}
		allButtons = append(allButtons, buttons)
	}

	// Print for debugging
	for i := 0; i < len(indicators); i++ {
		logger.Printf("%d : %v , %v", i, indicators[i], allButtons[i])
	}

	// Try every combination of 1 button
	// Try every combination of 2 buttons
	for i = 0; i < len(indicators); i++ {
		indicator := indicators[i]
		buttons := allButtons[i]
		var numButtons int
	iter:
		for j := 1; j <= len(buttons); j++ {
			// Try every combination of j number of buttons
			combinations := internal.Comb(len(buttons), j)
			for _, combination := range combinations {
				numButtons = len(combination)
				// Init result
				result := map[int]int{}
				for ii := range len(indicator) {
					result[ii] = 0
				}
				for _, c := range combination {
					button := buttons[c]
					for _, b := range button {
						result[b] = (result[b] + 1) % 2
					}
				}
				// logger.Printf("Checking combination: %v", combination)
				// logger.Printf("Result: %v", result)

				// Compare resulting indicator to required indicator
				if reflect.DeepEqual(result, indicator) {
					break iter
				}
			}
		}
		logger.Printf("%v : solved using %d buttons", indicator, numButtons)
		answer += numButtons
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

	joltages := []map[int]int{}
	allButtons := [][][]int{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		// Read buttons
		buttons := [][]int{}
		for _, button := range lineSplit[1 : len(lineSplit)-1] {
			b := []int{}
			for _, char := range button[1 : len(button)-1] {
				if char == ',' {
					continue
				}
				num, err := strconv.Atoi(string(char))
				if err != nil {
					logger.Fatalf("error converting string %s to integer: %v\n", string(char), err)
				}
				b = append(b, num)
			}
			buttons = append(buttons, b)
		}
		allButtons = append(allButtons, buttons)

		// Read joltages
		joltage := map[int]int{}
		j := strings.Split(lineSplit[len(lineSplit)-1][1:len(lineSplit[len(lineSplit)-1])-1], ",")
		for i := range j {
			num, err := strconv.Atoi(j[i])
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", j[i], err)
			}
			joltage[i] = num
		}
		joltages = append(joltages, joltage)
	}

	// Print for debugging
	for i := 0; i < len(joltages); i++ {
		logger.Printf("%d : Joltages: %v | Buttons: %v", i, joltages[i], allButtons[i])
	}

	// Try every combination of 1 button
	// Try every combination of 2 buttons
	// ...
	// Up until sum of all joltages
	for i = 0; i < len(joltages); i++ { // len(joltages)
		joltage := joltages[i]
		buttons := allButtons[i]
		joltageSum := 0
		for _, v := range joltage {
			joltageSum += v
		}
		var numButtons int
	iter:
		for j := 1; j <= joltageSum; j++ {
			// Try every combination of j number of buttons
			combinations := internal.CombWithReplacement(len(buttons), j)
			for _, combination := range combinations {
				numButtons = j
				// Init result
				result := map[int]int{}
				for ii := range len(joltage) {
					result[ii] = 0
				}
				for _, c := range combination {
					button := buttons[c]
					for _, b := range button {
						result[b] = result[b] + 1
					}
				}
				// logger.Printf("Checking combination: %v", combination)
				// logger.Printf("Result: %v", result)

				// Compare resulting joltage to required joltage
				if reflect.DeepEqual(result, joltage) {
					break iter
				}
			}
			if j == joltageSum {
				logger.Fatalf("no solution found for joltage: %v", joltage)
			}
		}
		logger.Printf("%v : solved using %d buttons", joltage, numButtons)
		answer += numButtons
	}

	return answer
}
