package day04

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day04/input.txt"
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
	grid := make([][]string, 0) // declare empty 2d array slice grid[y][x] or grid[r][c]

	// Read text into grid
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		row := make([]string, 0)
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			row = append(row, char)
		}
		grid = append(grid, row)

	}
	// logger.Printf("%v", grid)

	maxY := len(grid)
	maxX := len(grid[0])

	for y, row := range grid {
		for x, char := range row {
			// logger.Printf("(%d,%d) = %s", y, x, char)
			if char != "@" {
				continue
			}
			numNeighbouringToiletRolls := 0
			neighbours := [][]int{{x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}, {x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}}
			for _, n := range neighbours {
				n_x := n[0]
				n_y := n[1]
				// Skip if the neighbour is outside of the grid
				if n_x < 0 || n_x >= maxX || n_y < 0 || n_y >= maxY {
					continue
				}
				if grid[n_y][n_x] == "@" {
					numNeighbouringToiletRolls++
				}
			}
			logger.Printf("(%d,%d)=%s has %d neighbouring toilet rools", y, x, char, numNeighbouringToiletRolls)

			if numNeighbouringToiletRolls < 4 {
				answer++
			}
		}
	}

	return answer
}

func printGrid(grid [][]string, logger *log.Logger) {
	for _, row := range grid {
		logger.Println(row)
	}
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0
	grid := make([][]string, 0) // declare empty 2d array slice grid[y][x] or grid[r][c]

	// Read text into grid
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		row := make([]string, 0)
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			row = append(row, char)
		}
		grid = append(grid, row)

	}

	// logger.Printf("%v", grid)

	maxY := len(grid)
	maxX := len(grid[0])

	toiletRollsToRemove := make([][]int, 0)
	passes := 0
	for {
		passes++
		// Find out which toilet rolls we can remove (fewer than 4 neighbouring toilet rolls)
		for y, row := range grid {
			for x, char := range row {
				// logger.Printf("(%d,%d) = %s", y, x, char)
				if char != "@" {
					continue
				}
				numNeighbouringToiletRolls := 0
				neighbours := [][]int{{x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}, {x, y + 1}, {x - 1, y + 1}, {x - 1, y}, {x - 1, y - 1}}
				for _, n := range neighbours {
					n_x := n[0]
					n_y := n[1]
					// Skip if the neighbour is outside of the grid
					if n_x < 0 || n_x >= maxX || n_y < 0 || n_y >= maxY {
						continue
					}
					if grid[n_y][n_x] == "@" {
						numNeighbouringToiletRolls++
					}
				}
				// logger.Printf("(%d,%d)=%s has %d neighbouring toilet rools", y, x, char, numNeighbouringToiletRolls)
				if numNeighbouringToiletRolls < 4 {
					toiletRollsToRemove = append(toiletRollsToRemove, []int{x, y})
					answer++
				}
			}
		}

		// Remove all toilet rolls in toiletRollsToRemove from the grid
		for i := len(toiletRollsToRemove) - 1; i >= 0; i-- {
			x := toiletRollsToRemove[i][0]
			y := toiletRollsToRemove[i][1]
			grid[y][x] = "x"
		}

		// Break if we have found no toilet rolls to remove
		if len(toiletRollsToRemove) <= 0 {
			break
		}

		// Print the grid for debugging
		logger.Printf("--- Showing grid after %d passes ---", passes)
		logger.Printf("Removed %d toilet rolls", len(toiletRollsToRemove))
		printGrid(grid, logger)

		// Reset this array
		toiletRollsToRemove = make([][]int, 0) // reset the array
	}

	return answer
}
