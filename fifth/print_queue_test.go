package fifth

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrintQueue(t *testing.T) {
	rules_input := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}
	// input := []string{
	// 	"75,47,61,53,29",
	// 	"97,61,53,29,13",
	// 	"75,29,13",
	// 	"75,97,47,61,53",
	// 	"61,13,29",
	// 	"97,13,75,29,47",
	// }
	t.Run("Add Rules", func(t *testing.T) {
		ruleSet := NewRuleSet(rules_input)
		before := ruleSet.BeforeRules[53]
		after := ruleSet.AfterRules[53]
		if len(after) != 2 {
			t.Errorf("Expected 2, got %d", len(after))
		}
		if len(before) != 4 {
			t.Errorf("Expected 4, got %d", len(before))
		}

	})
	t.Run("Simple ruleset validation", func(t *testing.T) {
		test_sequence := []int{75, 47, 61, 53, 29}
		ruleSet := NewRuleSet(rules_input)
		got := ruleSet.IsUpdateValid(test_sequence)
		expect := true
		if got != expect {
			t.Errorf("Expected %v, got %v", expect, got)
		}
	})
	t.Run("Negative validation", func(t *testing.T) {
		test_sequence := []int{75, 97, 47, 61, 53}
		test_2_sequence := []int{61, 13, 29}
		ruleSet := NewRuleSet(rules_input)
		got := ruleSet.IsUpdateValid(test_sequence)
		got_2 := ruleSet.IsUpdateValid(test_2_sequence)
		expect := false
		if got != expect {
			t.Errorf("Expected %v, got %v", expect, got)
		}
		if got_2 != expect {
			t.Errorf("Expected %v, got %v", expect, got_2)
		}
	})
	t.Run("Correct the incorrect line", func(t *testing.T) {
		incorrect_line := []int{75, 97, 47, 61, 53}
		corrected_line := []int{97, 75, 47, 61, 53}
		ruleSet := NewRuleSet(rules_input)
		got := ruleSet.CorrectUpdate(incorrect_line)
		assert.DeepEqual(t, got, corrected_line)
	})
	t.Run("Correct the incorrect line 2", func(t *testing.T) {
		incorrect_line := []int{97, 13, 75, 29, 47}
		corrected_line := []int{97, 75, 47, 29, 13}
		ruleSet := NewRuleSet(rules_input)
		got := ruleSet.CorrectUpdate(incorrect_line)
		assert.DeepEqual(t, got, corrected_line)
	})
	t.Run("Full pass input on the ruleset", func(t *testing.T) {
		input := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}
		ruleSet := NewRuleSet(rules_input)
		valid_lines := make([][]int, 0)
		invalid_lines := make([][]int, 0)
		for _, update := range input {
			if ruleSet.IsUpdateValid(update) {
				valid_lines = append(valid_lines, update)
			} else {
				invalid_lines = append(invalid_lines, update)
			}
		}
		count := SumMiddleElements(valid_lines)
		corrected_lines := make([][]int, len(invalid_lines))
		for i, line := range invalid_lines {
			corrected_lines[i] = ruleSet.CorrectUpdate(line)
		}
		count_corrected := SumMiddleElements(corrected_lines)
		assert.Equal(t, count, 143)
		assert.Equal(t, count_corrected, 123)
	})
}
