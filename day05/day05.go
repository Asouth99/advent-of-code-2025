package day05

import (
	internal "aoc2025/internal/utils"
	"bufio"
	"errors"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day05/input.txt"
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
	freshIdRanges := []string{}
	ingredientIds := []int{}

	scanner := bufio.NewScanner(f)
	i := -1
	hasFinishedRanges := false
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if line == "" {
			hasFinishedRanges = true
			continue
		}
		if !hasFinishedRanges {
			freshIdRanges = append(freshIdRanges, line)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", line, err)
			}
			ingredientIds = append(ingredientIds, id)
		}
		// logger.Printf("%d : %s", i, line)
	}

	// Check each ingredient ID and see if it lies within any of the fresh ID ranges
	for _, ingredientId := range ingredientIds {
		logger.Printf("Checking ingredient ID %d", ingredientId)
		// Check against each range
		for _, freshIdRange := range freshIdRanges {
			idRange := strings.Split(freshIdRange, "-")
			min, err := strconv.Atoi(idRange[0])
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", idRange[0], err)
			}
			max, err := strconv.Atoi(idRange[1])
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", idRange[1], err)
			}

			if ingredientId <= max && ingredientId >= min {
				logger.Printf("Ingredient ID %d is fresh because it falls in range %s", ingredientId, idRange)
				answer++
				break
			}
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

	answer := 0
	freshIdRanges := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if line == "" {
			break
		}
		idRange := strings.Split(line, "-")
		min, err := strconv.Atoi(idRange[0])
		if err != nil {
			logger.Fatalf("error converting string %s to integer: %v\n", idRange[0], err)
		}
		max, err := strconv.Atoi(idRange[1])
		if err != nil {
			logger.Fatalf("error converting string %s to integer: %v\n", idRange[1], err)
		}
		freshIdRanges = append(freshIdRanges, []int{min, max})
	}

	// Sort ranges by their minimum
	// logger.Printf("Sorting %v", freshIdRanges)
	sort.Slice(freshIdRanges, func(i, j int) bool { return freshIdRanges[i][0] < freshIdRanges[j][0] })
	logger.Printf("Sorted ids: %v", freshIdRanges)

	// Combine ranges that overlap
	length := len(freshIdRanges)
	for i := 0; i < length-1; i++ {
		idRange := freshIdRanges[i]
		min := idRange[0]
		max := idRange[1]
		idRangeNext := freshIdRanges[i+1]
		minNext := idRangeNext[0]
		maxNext := idRangeNext[1]

		if minNext <= max {
			// The ranges overlap so combine the two ranges into min(min, minNext) - max(max, maxNext)
			minNew := internal.Min(min, minNext)
			maxNew := internal.Max(max, maxNext)
			freshIdRanges = slices.Delete(freshIdRanges, i+1, i+2)
			length--
			i--
			idRange[0] = minNew
			idRange[1] = maxNew
		}

	}
	logger.Printf("Combined ids: %v", freshIdRanges)

	// Calculate difference for every range now that we know there are no overlaps
	for _, freshIdRange := range freshIdRanges {
		min := freshIdRange[0]
		max := freshIdRange[1]
		answer += max - min + 1
	}

	return answer
}
