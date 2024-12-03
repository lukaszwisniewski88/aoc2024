package third

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ExtractMul(input string) (res [][]string) {
	exp := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := exp.FindAllStringSubmatch(input, -1)
	return matches
}
func ExtractDoSubstring(input string) string {
	var sections = struct {
		start []int
		end   []int
	}{
		start: []int{0},
		end:   []int{},
	}
	const (
		enabled  = 1
		disabled = 0
	)
	state := enabled
	for index, char := range input {
		if state == enabled {
			// search for don't
			if char == []rune("d")[0] {
				closing_index := strings.Index(input[index:], ")")
				if closing_index == -1 {
					break
				}
				substring := input[index : index+closing_index+1]
				if substring == "don't()" {
					state = disabled
					sections.end = append(sections.end, index)
				}
			}
		} else if state == disabled {
			// search for do
			if char == []rune("d")[0] {
				closing_index := strings.Index(input[index:], ")")
				if closing_index == -1 {
					break
				}
				substring := input[index : closing_index+index+1]
				if substring == "do()" {
					state = enabled
					sections.start = append(sections.start, index+closing_index+1)
				}
			}
		}
	}
	sections.end = append(sections.end, len(input))
	result := ""
	for i, section := range sections.start {
		result += input[section:sections.end[i]]
	}
	return result

}

func ExtractMulAndDo(input string) (res [][]string) {
	mul_exp := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for index, char := range input {
		switch char {
		case []rune("m")[0]:
			closing_index := strings.Index(input[index:], ")")
			if closing_index == 0 || closing_index <= index {
				break
			}
			substring := input[index:closing_index]
			if mul_exp.MatchString(substring) {
				matches := mul_exp.FindAllStringSubmatch(substring, -1)
				res = append(res, matches[0])
			}
		}
	}
	fmt.Println(res)
	return res
}

func FilterExtractedEnabledMul(input [][]string) [][]string {
	var result [][]string
	const (
		enabled  = 1
		disabled = 0
	)
	state := enabled
	for _, match := range input {
		if match[0] == "do()" {
			state = enabled

		} else if match[0] == "don't()" {
			state = disabled

		} else {
			if state == enabled {
				result = append(result, match)
			}
		}

	}
	return result
}

func AddExtractedMul(input [][]string) int {
	sum := 0
	for _, match := range input {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}
	return sum
}
