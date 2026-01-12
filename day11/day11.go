package day11

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day11/input.txt"
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

type device struct {
	value   string
	outputs []*device
}

func (this *device) countPaths() int {
	// fmt.Printf("%v has length %d\n", this.outputs, len(this.outputs))
	if len(this.outputs) == 0 {
		return 1
	}

	totalPaths := 0
	for _, d := range this.outputs {
		if d != nil {
			totalPaths += d.countPaths()
		}
	}
	return totalPaths
}

func (this *device) getPathsToOut() [][]string {
	if this.value == "out" {
		return [][]string{{this.value}}
	}

	allPaths := [][]string{}
	for _, d := range this.outputs {
		if d != nil {
			childPaths := d.getPathsToOut()
			for _, p := range childPaths {
				allPaths = append(allPaths, append([]string{this.value}, p...))
			}
		}
	}
	return allPaths
}

func (this *device) countPathsToOutwithDacFft(hasDac bool, hasFft bool) int {

	hasDac = hasDac || this.value == "dac"
	hasFft = hasFft || this.value == "fft"

	if this.value == "out" {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	count := 0
	for _, d := range this.outputs {
		count += d.countPathsToOutwithDacFft(hasDac, hasFft)
	}
	return count
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	devices := map[string]*device{}

	// Read input to create all devices
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		deviceValue := lineArr[0][:len(lineArr[0])-1] // Remove colon at the end
		devices[deviceValue] = &device{value: deviceValue, outputs: []*device{}}
	}
	// Read input again to add device outputs
	_, err = f.Seek(0, io.SeekStart) // Reset the file back to the start
	if err != nil {
		logger.Fatal(err)
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		deviceValue := lineArr[0][:len(lineArr[0])-1] // Remove colon at the end
		deviceOutputs := lineArr[1:]
		device := devices[deviceValue]
		for _, d := range deviceOutputs {
			if devices[d] != nil {
				device.outputs = append(device.outputs, devices[d])
			}
		}
	}

	// Print all devices for debugging
	for _, d := range devices {
		outputValues := []string{}
		for _, o := range d.outputs {
			if o != nil {
				outputValues = append(outputValues, o.value)
			}
		}
		logger.Printf("%s: [%s]", d.value, strings.Join(outputValues, ", "))
	}

	logger.Println("Counting number of paths from device 'you'")
	answer := devices["you"].countPaths()

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	devices := map[string]*device{}
	answer := 0

	// Read input to create all devices
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		deviceValue := lineArr[0][:len(lineArr[0])-1] // Remove colon at the end
		devices[deviceValue] = &device{value: deviceValue, outputs: []*device{}}
	}
	// Add device for 'out'
	devices["out"] = &device{value: "out", outputs: []*device{}}
	// Read input again to add device outputs
	_, err = f.Seek(0, io.SeekStart) // Reset the file back to the start
	if err != nil {
		logger.Fatal(err)
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, " ")
		deviceValue := lineArr[0][:len(lineArr[0])-1] // Remove colon at the end
		deviceOutputs := lineArr[1:]
		device := devices[deviceValue]
		for _, d := range deviceOutputs {
			if devices[d] != nil {
				device.outputs = append(device.outputs, devices[d])
			}
		}
	}

	// Print all devices for debugging
	for _, d := range devices {
		outputValues := []string{}
		for _, o := range d.outputs {
			if o != nil {
				outputValues = append(outputValues, o.value)
			}
		}
		logger.Printf("%s: [%s]", d.value, strings.Join(outputValues, ", "))
	}

	// Get all paths from svr to out
	// logger.Println("Getting all paths from 'svr' to 'out'")
	// paths := devices["svr"].getPathsToOut()
	// for _, p := range paths {
	// 	logger.Println(p)
	// 	hasFft := false
	// 	hasDac := false
	// 	for _, d := range p {
	// 		if d == "dac" {
	// 			hasDac = true
	// 		}
	// 		if d == "fft" {
	// 			hasFft = true
	// 		}
	// 	}
	// 	if hasDac && hasFft {
	// 		answer += 1
	// 	}
	// }

	// Count all paths from 'svr' to 'out' with 'dac' and 'fft'
	answer = devices["svr"].countPathsToOutwithDacFft(false, false)

	return answer
}
