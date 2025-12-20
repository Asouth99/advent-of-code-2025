package day07

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day07/input.txt"
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
	startCoord := []int{}

	// Read input into grid
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		row := make([]string, 0)
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if char == "S" {
				startCoord = []int{i, j}
			}
			row = append(row, char)
		}
		grid = append(grid, row)
	}

	maxY := len(grid)
	maxX := len(grid[0])

	// Debugging
	logger.Printf("Starting Coord: %v", startCoord)
	logger.Printf("Printing grid of %d by %d", maxX, maxY)
	for y, row := range grid {
		// for x, char := range row {
		// 	logger.Printf("(%d,%d) = %s", y, x, char)
		// }
		logger.Printf("%02d: %v", y, row)
	}

	// Perform iterations
	beamEnds := [][]int{startCoord}
	count := 0
	for len(beamEnds) > 0 {

		// Debugging
		logger.Printf("--- Printing grid after %d iterations ---", count)
		logger.Printf("BeamEnds: %v", beamEnds)

		// Loop through every beam end and perform iteration
		newBeamEnds := [][]int{}
		for _, beamEnd := range beamEnds {
			current := beamEnd
			y := current[0]
			x := current[1]

			// Check what is underneath the beam end
			// if "." Add a beamEnd underneath
			// if "^" Add a beamEnd to the BL and BR (if BL or BR is a ".")
			// if bottom of grid Do nothing
			if y >= len(grid)-1 {
				continue
			}
			switch grid[y+1][x] {
			case ".":
				newBeamEnds = append(newBeamEnds, []int{y + 1, x})
				grid[y+1][x] = "|"
			case "^":
				if (grid[y+1][x-1] == ".") || (grid[y+1][x+1] == ".") {
					answer++
				}
				if grid[y+1][x+1] == "." {
					newBeamEnds = append(newBeamEnds, []int{y + 1, x + 1})
					grid[y+1][x+1] = "|"
				}
				if grid[y+1][x-1] == "." {
					newBeamEnds = append(newBeamEnds, []int{y + 1, x - 1})
					grid[y+1][x-1] = "|"
				}
			default:
				continue
			}
		}
		count++

		// Empty the beamEnds array and reset it to the newBeamEnds
		beamEnds = newBeamEnds

		// Debugging. Print grid after every loop
		for y, row := range grid {
			logger.Printf("%02d: %v", y, row)
		}
		logger.Printf("Splits: %d", answer)
		logger.Printf("NewBeamEnds: %v", newBeamEnds)
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
	grid := make([][]string, 0) // declare empty 2d array slice grid[y][x] or grid[r][c]
	startCoord := []int{}

	// Read input into grid
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		row := make([]string, 0)
		for j := 0; j < len(line); j++ {
			char := string(line[j])
			if char == "S" {
				startCoord = []int{i, j}
			}
			row = append(row, char)
		}
		grid = append(grid, row)
	}

	maxY := len(grid)
	maxX := len(grid[0])

	// Debugging
	logger.Printf("Starting Coord: %v", startCoord)
	logger.Printf("Printing grid of %d by %d", maxX, maxY)
	for y, row := range grid {
		logger.Printf("%02d: %v", y, row)
	}

	// Store the number of beams in a column
	numBeamsInColumns := map[int]int{startCoord[1]: 1} // There is 1 beam in the starting column

	// Loop through line by line because the beams move from top down
	for _, row := range grid {
		for x, char := range row {
			beamsInColumn := numBeamsInColumns[x]
			if char == "^" && x > 0 && x < maxX {
				numBeamsInColumns[x-1] += beamsInColumn
				numBeamsInColumns[x+1] += beamsInColumn
				delete(numBeamsInColumns, x)
			}
		}
	}

	// Sum up the values in numBeamsInColumn
	for _, beamsInColumn := range numBeamsInColumns {
		answer += beamsInColumn
	}

	return answer
}
