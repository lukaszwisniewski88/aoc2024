package fifth

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ProcessDayFive(input string) (string, error) {
	file, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()
	input_text, err := io.ReadAll(file)
	if err != nil {
		return "Error reading file", err
	}
	// empty line is the separator of the rules and the input
	splitted := strings.Split(string(input_text), "\n\n")
	rules := splitted[0]
	updates := splitted[1]
	valid_lines := make([][]int, 0)
	ruleSet := NewRuleSet(strings.Split(rules, "\n"))
	updates_list := strings.Split(updates, "\n")
	for _, update := range updates_list {
		update_list := strings.Split(update, ",")
		update_int := make([]int, len(update_list))
		for i, v := range update_list {
			update_int[i], err = strconv.Atoi(v)
			if err != nil {
				continue
			}
		}
		if ruleSet.IsUpdateValid(update_int) {
			valid_lines = append(valid_lines, update_int)
		}
	}
	count := SumMiddleElements(valid_lines)
	return fmt.Sprintf("%d", count), nil
}

func SumMiddleElements(lines [][]int) int {
	sum := 0
	for _, line := range lines {
		sum += line[len(line)/2]
	}
	return sum
}
