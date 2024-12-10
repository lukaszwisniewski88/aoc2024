package fifth

import (
	"fmt"
)

type Rule struct {
	Before int
	After  int
}

type RuleSet struct {
	BeforeRules map[int][]int
	AfterRules  map[int][]int
}

func NewRuleSet(rules []string) RuleSet {
	ruleSet := RuleSet{
		BeforeRules: make(map[int][]int),
		AfterRules:  make(map[int][]int),
	}
	for _, rule := range rules {
		ruleSet.addRule(rule)
	}
	return ruleSet
}

func (rs *RuleSet) addRule(rule string) {
	before, after := parseRule(rule)
	rs.AfterRules[before] = append(rs.AfterRules[before], after)
	rs.BeforeRules[after] = append(rs.BeforeRules[after], before)
}

func parseRule(rule string) (int, int) {
	var before, after int
	fmt.Sscanf(rule, "%d|%d", &before, &after)
	return before, after
}
func (s *RuleSet) canBeBefore(before, after int) bool {
	// no vialating any rule
	after_rules, ok := s.AfterRules[after]
	if !ok {
		return true
	}
	for _, rule := range after_rules {
		if rule == before {
			return false
		}
	}
	return true
}
func (s *RuleSet) canBeAfter(after, before int) bool {
	// no vialating any rule
	before_rules, ok := s.BeforeRules[before]
	if !ok {
		return true
	}
	for _, rule := range before_rules {
		if rule == after {
			return false
		}
	}
	return true
}

func (s *RuleSet) GetBeforeRules(key int) []int {
	return s.BeforeRules[key]
}

func (s *RuleSet) IsUpdateValid(line []int) bool {
	for i, key := range line {
		key_is_before := line[i+1:]
		key_is_after := line[:i]
		if len(key_is_before) > 0 {
			for _, k := range key_is_before {
				if !s.canBeBefore(key, k) {
					return false
				}
			}
		}
		if len(key_is_after) > 0 {
			for _, k := range key_is_after {
				if !s.canBeAfter(key, k) {
					return false
				}
			}
		}
	}
	return true
}
func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
func (s *RuleSet) CanBeBefore(number int, beforeSlice []int) (bool, int) {
	for _, after := range beforeSlice {
		if !s.canBeAfter(number, after) {
			return false, after
		}
	}
	return true, -1
}
func (s *RuleSet) CanBeAfter(number int, beforeSlice []int) (bool, int) {
	for _, before := range beforeSlice {
		if !s.canBeBefore(number, before) {
			return false, before
		}
	}
	return true, -1
}

func (s *RuleSet) CorrectUpdate(line []int) []int {
	for !s.IsUpdateValid(line) {
		for i, key := range line {
			after_slice := line[i+1:]
			before_slice := line[:i]
			if len(after_slice) > 0 {
				can_be_after, before := s.CanBeAfter(key, after_slice)
				if !can_be_after {
					index := indexOf(before, line)
					line[i], line[index] = line[index], line[i]

				}
			}
			if len(before_slice) > 0 {
				can_be_before, after := s.CanBeBefore(key, before_slice)
				if !can_be_before {
					index := indexOf(after, line)
					line[i], line[index] = line[index], line[i]
				}
			}
		}
	}

	return line
}
