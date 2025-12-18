package day02

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (int, error) {
	file := "./day02/input.txt"
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

func isInvalidIdPart1(id int) bool {
	// Invalid IDs are any ID which is made only of some sequence of digits repeated twice.

	stringId := strconv.Itoa(id)
	if len(stringId)%2 != 0 { // Must be valid if there is an odd number of digits
		return false
	}

	// Split string in half and see if firstHalf = secondHalf
	firstHalf := stringId[:len(stringId)/2]
	secondHalf := stringId[len(stringId)/2:]
	return firstHalf == secondHalf
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
		line := scanner.Text() // Input is a single long line
		// logger.Printf("%d : %s", i, line)

		idRanges := strings.Split(line, ",") // IDs are comma separated

		for _, idRange := range idRanges {
			logger.Println(idRange)
			ids := strings.Split(idRange, "-") // First ID and Last ID are separated by a dash
			firstId, err := strconv.Atoi(ids[0])
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", ids[0], err)
			}
			lastId, err := strconv.Atoi(ids[1])
			if err != nil {
				logger.Fatalf("error converting string %s to integer: %v\n", ids[1], err)
			}

			// Check if any ID between firstId and lastId inclusive is invalid. If it is then add it to the answer
			for id := firstId; id <= lastId; id++ {
				// logger.Println(id)
				if isInvalidIdPart1(id) {
					logger.Printf("%d is invalid. Updating answer", id)
					answer += id
				}
			}
		}

	}
	return answer
}

func isInvalidIdPart2(id int, logger *log.Logger) bool {
	// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice.
	// Examples:
	// 12341234 (1234 two times)
	// 123123123 (123 three times)
	// 1212121212 (12 five times)
	// 1111111 (1 seven times)
	// are all invalid IDs.

	// Patterns:
	// - 1 char repeated n number of times
	// 11   * a = aa
	// 111  * a = aaa
	// 1111 * a = aaaa
	// - 2 chars repeated n number of times
	// 101     * ab = abab
	// 10101   * ab = ababab
	// 1010101 * ab = abababab
	// - 3 chars repeated n number of times
	// 1001       * abc = abcabc
	// 1001001    * abc = abcabcabc
	// 1001001001 * abc = abcabcabcabc
	// - 4 chars repeated n number of times
	// 10001         * abcd = abcdabcd
	// 100010001     * abcd = abcdabcdabcd
	// 1000100010001 * abcd = abcdabcdabcdabcd

	// convert to string cos its useful
	stringId := strconv.Itoa(id)

	// testing i chars repeated n times
	for i := 1; i <= len(stringId)/2; i++ {
		// Continue if string is not divisible by i. Won't be a full number of repetitions
		if len(stringId)%i != 0 {
			continue
		}

		// i chars repeated repetitions times
		repetitions := len(stringId) / i

		logger.Printf("Checking %d for %d chars repeated %d times", id, i, repetitions)

		numZeros := strings.Repeat("0", i-1)
		divisorString := "1" + strings.Repeat(numZeros+"1", repetitions-1) // generates string like 1001 or 1001001
		divisor, err := strconv.Atoi(divisorString)
		if err != nil {
			logger.Fatalf("error converting string %s to integer: %v\n", divisorString, err)
		}
		if id%divisor == 0 {
			resultString := strconv.Itoa(id / divisor)
			if len(resultString) == i {
				logger.Printf("%d is divisible by %d resulting in %s\n", id, divisor, resultString)
				return true
			}
		}
	}

	return false

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
		line := scanner.Text()               // Input is a single long line
		idRanges := strings.Split(line, ",") // IDs are comma separated

		for _, idRange := range idRanges {
			logger.Println(idRange)
			ids := strings.Split(idRange, "-") // First ID and Last ID are separated by a dash
			firstId, err := strconv.Atoi(ids[0])
			if err != nil {
				logger.Fatalf("error converting string %d to integer: %v\n", firstId, err)
			}
			lastId, err := strconv.Atoi(ids[1])
			if err != nil {
				logger.Fatalf("error converting string %d to integer: %v\n", lastId, err)
			}

			// Check if any ID between firstId and lastId inclusive is invalid. If it is then add it to the answer
			for id := firstId; id <= lastId; id++ {
				// logger.Println(id)
				if isInvalidIdPart2(id, logger) {
					logger.Printf("%d is invalid. Updating answer", id)
					answer += id
				}
			}
		}
	}
	return answer
}
