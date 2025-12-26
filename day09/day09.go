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

	type edge struct {
		from, to *point
	}
	verticalEdges := []edge{}
	horizontalEdges := []edge{}
	// Coords in input are given in order going clockwise around the shape
	for i := range points {
		p1, p2 := points[i], points[(i+1)%len(points)]
		if p1.x == p2.x {
			verticalEdges = append(verticalEdges, edge{from: &p1, to: &p2})
		} else if p1.y == p2.y {
			horizontalEdges = append(horizontalEdges, edge{from: &p1, to: &p2})
		} else {
			logger.Fatalln("points are neither vertically or horizontally aligned")
		}
	}

	// Print for debugging
	for i, p := range points {
		logger.Printf("%d: (%.0f,%.0f)", i, p.x, p.y)
	}

	// Get area of rectangles of each point to every other point
	var maximumArea float64 = 0
	for i, p1 := range points {
	points:
		for _, p2 := range points[i+1:] {
			area := (math.Abs(p1.x-p2.x) + 1) * (math.Abs(p1.y-p2.y) + 1)

			// No need to check the area if it is smaller than the largest so far
			if area < maximumArea {
				continue
			}

			topEdge := edge{from: &point{x: min(p1.x, p2.x), y: min(p1.y, p2.y)}, to: &point{x: max(p1.x, p2.x), y: min(p1.y, p2.y)}}
			rightEdge := edge{from: &point{x: max(p1.x, p2.x), y: min(p1.y, p2.y)}, to: &point{x: max(p1.x, p2.x), y: max(p1.y, p2.y)}}
			bottomEdge := edge{from: &point{x: max(p1.x, p2.x), y: max(p1.y, p2.y)}, to: &point{x: min(p1.x, p2.x), y: max(p1.y, p2.y)}}
			leftEdge := edge{from: &point{x: min(p1.x, p2.x), y: max(p1.y, p2.y)}, to: &point{x: min(p1.x, p2.x), y: min(p1.y, p2.y)}}
			logger.Printf("Top: (%.0f,%.0f)-(%.0f,%.0f)  | Right: (%.0f,%.0f)-(%.0f,%.0f) | Bottom: (%.0f,%.0f)-(%.0f,%.0f) | Left: (%.0f,%.0f)-(%.0f,%.0f)", topEdge.from.x, topEdge.from.y, topEdge.to.x, topEdge.to.y, rightEdge.from.x, rightEdge.from.y, rightEdge.to.x, rightEdge.to.y, bottomEdge.from.x, bottomEdge.from.y, bottomEdge.to.x, bottomEdge.to.y, leftEdge.from.x, leftEdge.from.y, leftEdge.to.x, leftEdge.to.y)

			// Check if all edges of the rectangle lie within the horizontal and vertical edges of the input
			for _, edge := range horizontalEdges {
				if edge.from.x < edge.to.x {
					// We are a top edge of the input shape
					// Check if either of the rectangle vertical edges are between the horizontal edge
					if ((edge.from.x <= leftEdge.from.x && edge.to.x > leftEdge.from.x) || (edge.from.x < rightEdge.from.x && edge.to.x >= rightEdge.from.x)) && leftEdge.from.y != rightEdge.from.y && edge.from.y > topEdge.from.y {
						logger.Printf("Top edge %v-%v is above top edge %v-%v of input shape", *topEdge.from, *topEdge.to, *edge.from, *edge.to)
						continue points
					}
				} else {
					// We are a bottom edge of the input shape
					// Check if either of the rectangle vertical edges are between the horizontal edge
					if ((edge.to.x <= leftEdge.from.x && edge.from.x > leftEdge.from.x) || (edge.to.x < rightEdge.from.x && edge.from.x >= rightEdge.from.x)) && leftEdge.from.y != rightEdge.from.y && edge.from.y < bottomEdge.from.y {
						logger.Printf("Bottom edge %v-%v is below bottom edge %v-%v of input shape", *bottomEdge.from, *bottomEdge.to, *edge.from, *edge.to)
						continue points
					}
				}
			}
			for _, edge := range verticalEdges {
				if edge.from.y < edge.to.y {
					// We are a right edge of the input shape
					// Check if either of the rectangle horizontal edges are between the vertical edge
					if ((edge.from.y <= topEdge.from.y && edge.to.y > topEdge.from.y) || (edge.from.y < bottomEdge.from.y && edge.to.y >= bottomEdge.from.y)) && bottomEdge.from.y != topEdge.from.y && edge.from.x < rightEdge.from.x {
						logger.Printf("Right edge %v-%v is right of right edge %v-%v of input shape", *rightEdge.from, *rightEdge.to, *edge.from, *edge.to)
						continue points
					}
				} else {
					// We are a left edge of the input shape
					// Check if either of the rectangle horizontal edges are between the vertical edge
					if ((edge.to.y <= topEdge.from.y && edge.from.y > topEdge.from.y) || (edge.to.y < bottomEdge.from.y && edge.from.y >= bottomEdge.from.y)) && bottomEdge.from.y != topEdge.from.y && edge.from.x > leftEdge.from.x {
						logger.Printf("Left edge %v-%v is left of left edge %v-%v of input shape", *leftEdge.from, *leftEdge.to, *edge.from, *edge.to)
						continue points
					}
				}
			}

			// Update the maximum if rectangle is OK
			if area > maximumArea {
				maximumArea = area
				logger.Printf("Rectangle area from %v to %v = %.0f", p1, p2, area)
			}
		}
	}

	logger.Printf("Max area possible is %.0f", maximumArea)
	answer := int(maximumArea)
	return answer
}

// 88983048 -> answer too low
