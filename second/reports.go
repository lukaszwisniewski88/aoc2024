package second

import (
	"bufio"
	"fmt"
	"lukaszwisniewski88/aoc2024/first"
	"math"
	"strconv"
	"strings"
)

func IsReportSafe(line []int) bool {
	const (
		INCREASING = iota << 1
		DECREASING
		EQUAL
		INDETERMINATE
	)
	var state int = INDETERMINATE
	for i, num := range line[:len(line)-1] {
		diff := num - line[i+1]
		if diff == 0 {
			return false
		}
		if math.Abs(float64(diff)) > 3 {
			return false
		}
		if diff > 0 {
			if state == DECREASING {
				return false
			}
			state = INCREASING
		} else {
			if state == INCREASING {
				return false
			}
			state = DECREASING
		}
	}
	return true
}
func RemoveIndex(s []int, index int) []int {
	newSlice := make([]int, 0)
	for i, num := range s {
		if i == index {
			continue
		}
		newSlice = append(newSlice, num)
	}
	return newSlice
}

func CanBeSafeByRemoving(line []int) (isSafe bool) {
	isSafe = false
	if IsReportSafe(line) {
		isSafe = true
		return isSafe
	}
	for i := range line {
		newLine := RemoveIndex(line, i)
		if IsReportSafe(newLine) {
			isSafe = true
			break
		}
	}

	return isSafe
}
func CountLessSafeReports(lines [][]int) int {
	count := 0
	for _, line := range lines {
		if CanBeSafeByRemoving(line) {
			count++
		}
	}
	return count
}

func CountSafeReports(lines [][]int) int {
	count := 0
	for _, line := range lines {
		if IsReportSafe(line) {
			count++
		}
	}
	return count
}

func GetInputLines(scanner *bufio.Scanner) ([][]int, error) {
	var result = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Fields(line)
		var lineInts = make([]int, 0)
		for _, s := range splitted {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			lineInts = append(lineInts, num)
		}
		result = append(result, lineInts)

	}
	return result, nil
}

func ProcessDayTwo(path string) (string, error) {
	scanner, err := first.OpenFile(path)
	if err != nil {
		return "", err
	}
	lines, err := GetInputLines(scanner)
	if err != nil {
		return "", err
	}
	safe := CountSafeReports(lines)
	lessSafe := CountLessSafeReports(lines)
	return fmt.Sprintf("Safe: %d, Less safe: %d", safe, lessSafe), nil
}
