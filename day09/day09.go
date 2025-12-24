package day09

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day09/input.txt"
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

type point struct {
	x, y float64
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read input file
	points := []point{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		nums := []float64{}
		for _, str := range strings.Split(line, ",") {
			num, err := strconv.Atoi(str)
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", str, err)
			}
			nums = append(nums, float64(num))
		}
		points = append(points, point{x: nums[0], y: nums[1]})
	}

	// Print for debugging
	for i, p := range points {
		logger.Printf("%d: (%.0f,%.0f)", i, p.x, p.y)
	}

	// Get distances of each point
	rectangles := map[*point]map[*point]float64{}
	for i, p1 := range points {
		rectangles[&p1] = map[*point]float64{}
		for _, p2 := range points[i+1:] {
			rectangles[&p1][&p2] = (math.Abs(p1.x-p2.x) + 1) * (math.Abs(p1.y-p2.y) + 1)
		}
	}

	// Get maximum distance
	var maximum float64 = 0
	var maxP1 point = point{}
	var maxP2 point = point{}
	for p1, p1Map := range rectangles {
		for p2, d := range p1Map {
			if d > maximum {
				maximum = d
				maxP1 = *p1
				maxP2 = *p2
			}
			logger.Printf("%v to %v = %.0f", *p1, *p2, d)
		}
	}
	logger.Printf("Max area %.0f from %v to %v", maximum, maxP1, maxP2)
	answer := int(maximum)
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
