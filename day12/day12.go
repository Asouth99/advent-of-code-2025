package day12

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day12/input.txt"
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

type shape struct {
	positions [][]rune
	area      int
}
type region struct {
	width  int
	height int
	shapes []int
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0
	allShapes := []shape{}
	allRegions := []region{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 2 { // Start of a new shape
			newShape := &shape{}
			newShape.area = 0
			for range 3 {
				if scanner.Scan() {
					lineArr := []rune{}
					for _, char := range scanner.Text() {
						if char == '#' {
							newShape.area++
						}
						lineArr = append(lineArr, char)
					}
					newShape.positions = append(newShape.positions, lineArr)
				}
			}
			scanner.Scan() // New line
			allShapes = append(allShapes, *newShape)
		} else if len(scanner.Bytes()) > 3 { // Start of a new region
			newRegion := &region{}
			line := scanner.Text()
			lineArr := strings.Split(line, ":")
			widthHeight := strings.Split(lineArr[0], "x")
			width, err := strconv.Atoi(widthHeight[0])
			if err != nil {
				logger.Fatalf("error converting string %s to int: %v\n", widthHeight[0], err)
			}
			height, err := strconv.Atoi(widthHeight[1])
			if err != nil {
				logger.Fatalf("error converting string %s to int: %v\n", widthHeight[1], err)
			}
			newRegion.width = width
			newRegion.height = height

			logger.Printf("<%v>", lineArr[1])
			for _, numStr := range strings.Split(strings.Trim(lineArr[1], " "), " ") {
				logger.Println(numStr)
				num, err := strconv.Atoi(numStr)
				if err != nil {
					logger.Fatalf("error converting string %s to int: %v\n", numStr, err)
				}
				newRegion.shapes = append(newRegion.shapes, num)
			}
			allRegions = append(allRegions, *newRegion)
		}

	}

	// Print all shapes
	for i := range allShapes {
		fmt.Printf("Shape %d - total area %d\n", i, allShapes[i].area)
		for _, row := range allShapes[i].positions {
			for _, char := range row {
				fmt.Printf("%c", char)
			}
			fmt.Println()
		}
		fmt.Println()
	}
	// Print all regions
	for i, r := range allRegions {
		logger.Printf("Region %d: %d x %d | Shapes: %v", i, r.width, r.height, r.shapes)
	}

	// Check each region
	for i, r := range allRegions {
		logger.Printf("Checking region %d\n", i)

		// Check how many 3by3 shapes can fit
		shapesCapacity := r.width / 3 * r.height / 3
		numShapes := 0
		for _, n := range r.shapes {
			numShapes += n
		}
		if numShapes > shapesCapacity {
			logger.Printf("  Region can hold %d loosely packed shapes. Total shapes count is %d\n", shapesCapacity, numShapes)
			continue
		}

		// Check if the areas fit
		regionArea := r.width * r.height
		totalShapeArea := 0

		for i, n := range r.shapes {
			totalShapeArea += allShapes[i].area * n
		}

		if regionArea < totalShapeArea {
			logger.Printf("  Regions area %d is less than all shapes area %d\n", regionArea, totalShapeArea)
		} else {
			logger.Printf("  Regions area %d is more than all shapes area %d\n", regionArea, totalShapeArea)
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
