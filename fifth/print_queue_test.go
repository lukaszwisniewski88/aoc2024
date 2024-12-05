package fifth

import "testing"

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
}
