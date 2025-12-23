package day08

import (
	internal "aoc2025/internal/utils"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day08/input.txt"
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

	var NUM_SHORT_CONNECTIONS int
	if inputFile == "./day08/input.txt" {
		NUM_SHORT_CONNECTIONS = 1000
	} else {
		NUM_SHORT_CONNECTIONS = 10
	}

	// Read input file
	junctionBoxes := [][]float64{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		arr := []float64{}
		for _, str := range strings.Split(line, ",") {
			num, err := strconv.Atoi(str)
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", str, err)
			}
			arr = append(arr, float64(num))
		}
		junctionBoxes = append(junctionBoxes, arr)
	}

	// Print for debugging
	// for i, box := range junctionBoxes {
	// 	logger.Printf("%02d : %v", i, box)
	// }

	// Find distances between each coord
	// var distanceMatrix [][]float64 // distanceMatrix[i][j] is distance from i to j. Has to obey i < j
	distanceMatrix := map[string]float64{} // distanceMatrix[i][j] is distance from i to j {[x1,y1,z1]-[x2,y2,z2] : float}
	for i, box := range junctionBoxes {
		// distances := []float64{}
		for j := i + 1; j < len(junctionBoxes); j++ {
			// distances = append(distances, internal.Distance(box, junctionBoxes[j]))
			key := fmt.Sprintf("%v-%v", box, junctionBoxes[j])
			distanceMatrix[key] = internal.Distance(box, junctionBoxes[j])
		}
		// distanceMatrix = append(distanceMatrix, distances)
	}

	// logger.Println(internal.Distance([]float64{906, 360, 560}, []float64{805, 96, 715}))

	// Print for debugging
	// logger.Println("Distance Matrix:")
	// for key, distances := range distanceMatrix {
	// 	logger.Printf("%s : %.0f", key, distances)
	// }

	// Define circuits
	circuits := []map[string]bool{} // [{"x0,y0,z0":true, "x1,y1,z1":true}, {"x2,y2,z2":true, "x3,y3,z3":true}]
	for numMinimums := 0; numMinimums < NUM_SHORT_CONNECTIONS; numMinimums++ {
		// Get smallest distance
		var minimum float64 = math.Inf(1)
		minKey := ""
		for key, distance := range distanceMatrix {
			if distance < minimum && distance != 0 {
				minimum = distance
				minKey = key
			}
		}
		delete(distanceMatrix, minKey) // Delete minimum from map so we dont grab it again
		coords := strings.Split(minKey, "-")
		logger.Printf("%d : Minmum distance is %f, connecting %s and %s", numMinimums, minimum, coords[0], coords[1])
		// Add it to a circuit
		var ok1, ok2 bool
		for i, circuit := range circuits {
			_, ok1 = circuit[coords[0]]
			_, ok2 = circuit[coords[1]]
			if ok1 || ok2 {
				circuit[coords[0]] = true
				circuit[coords[1]] = true
				logger.Printf("Added %s and %s to circuit %v", coords[0], coords[1], circuit)
				// break // Can't break because we need to check if we need to join any circuits together
				for j := i + 1; j < len(circuits); j++ {
					current := circuits[j]
					_, ok1 = current[coords[0]]
					_, ok2 = current[coords[1]]
					if ok1 || ok2 {
						// Join the two circuits
						for k, v := range current {
							circuit[k] = v
							delete(current, k) // Delete second circuit
						}
					}
				}
			}
		}
		if !ok1 && !ok2 {
			circuits = append(circuits, map[string]bool{coords[0]: true, coords[1]: true})
			logger.Printf("Creating a new circuit connecting %s and %s", coords[0], coords[1])
		}

	}

	// Print for debugging
	sort.Slice(circuits, func(a, b int) bool { return len(circuits[a]) > len(circuits[b]) })
	for i, circuit := range circuits {
		logger.Printf("%d : %d -> %v", i, len(circuit), circuit)
	}

	answer := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return answer
}

// Types required for part 2
type junctionBox struct {
	x, y, z float64
	circuit *circuit
}
type circuit struct {
	junctionBoxes []*junctionBox
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read input file
	scanner := bufio.NewScanner(f)
	i := -1
	allBoxes := []*junctionBox{}
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
		allBoxes = append(allBoxes, &junctionBox{
			x: nums[0],
			y: nums[1],
			z: nums[2],
		})
	}

	// Print for debugging
	for i, box := range allBoxes {
		logger.Printf("Box %02d : (%.0f, %.0f, %.0f)", i, box.x, box.y, box.z)
	}

	// Create distance matrix
	type connection struct {
		box1, box2 *junctionBox
		distance   float64
	}
	allConnections := []connection{}
	for i, box1 := range allBoxes {
		for _, box2 := range allBoxes[i+1:] {
			allConnections = append(allConnections, connection{
				box1:     box1,
				box2:     box2,
				distance: internal.Distance([]float64{box1.x, box1.y, box1.z}, []float64{box2.x, box2.y, box2.z}),
			})
		}
	}

	// Sort all of the distances
	sort.Slice(allConnections, func(i, j int) bool { return allConnections[i].distance < allConnections[j].distance })

	// Print for debugging
	// for _, connection := range allConnections {
	// 	logger.Printf("Distance between (%.0f, %.0f, %.0f) and (%.0f, %.0f, %.0f) is %.0f", connection.box1.x, connection.box1.y, connection.box1.z, connection.box2.x, connection.box2.y, connection.box2.z, connection.distance)
	// }

	// Loop through all connections and add the boxes to circuits
	allCircuits := []*circuit{}
	for len(allConnections) > 0 {
		shortestConnection := allConnections[0]
		allConnections = allConnections[1:]

		// Print for debugging
		// if shortestConnection.box1.circuit == nil {
		// 	logger.Printf("Box1 (%.0f, %.0f, %.0f) is part of no circuits", shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z)
		// } else {
		// 	logger.Printf("Box1 (%.0f, %.0f, %.0f) is part of a circuit with length %d", shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z, len(shortestConnection.box1.circuit.junctionBoxes))
		// }
		// if shortestConnection.box2.circuit == nil {
		// 	logger.Printf("Box2 (%.0f, %.0f, %.0f) is part of no circuits", shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z)
		// } else {
		// 	logger.Printf("Box2 (%.0f, %.0f, %.0f) is part of a circuit with length %d", shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z, len(shortestConnection.box2.circuit.junctionBoxes))
		// }

		// Check the circuits of each box
		if shortestConnection.box1.circuit == nil && shortestConnection.box2.circuit == nil {
			logger.Printf("Creating a new circuit for (%.0f, %.0f, %.0f) and (%.0f, %.0f, %.0f)", shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z, shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z)
			// Create a new circuit
			allCircuits = append(allCircuits, &circuit{junctionBoxes: []*junctionBox{shortestConnection.box1, shortestConnection.box2}})
			shortestConnection.box1.circuit = allCircuits[len(allCircuits)-1]
			shortestConnection.box2.circuit = allCircuits[len(allCircuits)-1]
		} else if shortestConnection.box1.circuit != nil && shortestConnection.box2.circuit == nil {
			logger.Printf("Adding (%.0f, %.0f, %.0f) to circuit of (%.0f, %.0f, %.0f)", shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z, shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z)
			// Add box2 to circuit of box1
			shortestConnection.box1.circuit.junctionBoxes = append(shortestConnection.box1.circuit.junctionBoxes, shortestConnection.box2)
			shortestConnection.box2.circuit = shortestConnection.box1.circuit
			// logger.Println(*shortestConnection.box2.circuit)
		} else if shortestConnection.box1.circuit == nil && shortestConnection.box2.circuit != nil {
			logger.Printf("Adding (%.0f, %.0f, %.0f) to circuit of (%.0f, %.0f, %.0f)", shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z, shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z)
			// Add box1 to circuit of box2
			shortestConnection.box2.circuit.junctionBoxes = append(shortestConnection.box2.circuit.junctionBoxes, shortestConnection.box1)
			shortestConnection.box1.circuit = shortestConnection.box2.circuit
		} else if shortestConnection.box1.circuit != nil && shortestConnection.box2.circuit != nil && shortestConnection.box1.circuit != shortestConnection.box2.circuit {
			logger.Printf("Merging circuits of (%.0f, %.0f, %.0f) and (%.0f, %.0f, %.0f)", shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z, shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z)
			// Merge circuits of box1 and box2
			oldCircuit := shortestConnection.box2.circuit
			for _, b := range shortestConnection.box2.circuit.junctionBoxes {
				shortestConnection.box1.circuit.junctionBoxes = append(shortestConnection.box1.circuit.junctionBoxes, b)
				b.circuit = shortestConnection.box1.circuit
			}
			for i, circuit := range allCircuits {
				if circuit == oldCircuit {
					allCircuits[len(allCircuits)-1], allCircuits[i] = allCircuits[i], allCircuits[len(allCircuits)-1] // move allCircuits[i] to the end
					allCircuits = allCircuits[0 : len(allCircuits)-1]                                                 // Remove the last circuit
					break
				}
			}
		}

		// Check if we have combined into a single circuit
		if len(allCircuits) == 1 && len(allCircuits[0].junctionBoxes) == len(allBoxes) {
			answer := shortestConnection.box1.x * shortestConnection.box2.x
			logger.Printf("Box (%.0f, %.0f, %.0f) and (%.0f, %.0f, %.0f) caused all the circuits to merge into one!", shortestConnection.box1.x, shortestConnection.box1.y, shortestConnection.box1.z, shortestConnection.box2.x, shortestConnection.box2.y, shortestConnection.box2.z)
			return int(answer)
		}

	}

	// Print for debugging
	for i, circuit := range allCircuits {
		logger.Printf("Circuit %d has %d junction boxes", i, len(circuit.junctionBoxes))
	}
	logger.Fatalln("should have returned in the for loop")

	return 0
}
