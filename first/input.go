package first

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OpenFile(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(file), nil
}

func GetInputLines(scanner *bufio.Scanner) ([][]int, error) {
	var result = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Fields(line)
		left, err := strconv.Atoi(splitted[0])
		if err != nil {
			return nil, err
		}
		right, err := strconv.Atoi(splitted[1])
		if err != nil {
			return nil, err
		}
		result = append(result, []int{left, right})
	}
	return result, nil
}

func ProcessDayOne(path string) (string, error) {
	scanner, err := OpenFile(path)
	if err != nil {
		return "", err
	}
	input, err := GetInputLines(scanner)
	if err != nil {
		return "", err
	}
	pairs, err := ProcessLines(input)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d, %d", pairs.GetSumDiffs(), pairs.GetSimilarityScore()), nil
}

func ProcessLines(input [][]int) (Pairs, error) {
	pairs := NewPairs()
	for _, line := range input {
		pairs.AddPair(line[0], line[1])
	}
	return *pairs, nil

}
