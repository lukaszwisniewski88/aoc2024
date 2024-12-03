package third

import (
	"testing"

	"gotest.tools/assert"
)

func TestMul(t *testing.T) {
	t.Run("Should read instructions of the test string", func(t *testing.T) {
		test_string := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
		expected := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
		result := ExtractMul(test_string)
		instructions := make([]string, 0)
		for _, match := range result {
			instructions = append(instructions, match[0])
		}
		assert.DeepEqual(t, expected, instructions)
	})
	t.Run("should sum multiplication of mul instructions", func(t *testing.T) {
		test_string := "xmul(2,4)%&mul[3,7]mul(4*!@^do_not_mul(5,5)+mul(32,64]mul(6,9!then(mul(11,8)mul(8,5))?(12,34)"
		expected := 2*4 + 5*5 + 11*8 + 8*5
		extracted := ExtractMul(test_string)
		result := AddExtractedMul(extracted)
		assert.Equal(t, expected, result)
	})
	t.Run("should build a string where only do() instructions are enabled", func(t *testing.T) {
		test_string := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		expected := "xmul(2,4)&mul[3,7]!^?mul(8,5))"
		result := ExtractDoSubstring(test_string)
		assert.Equal(t, expected, result)
	})
	t.Run("should sum after enabled filter mul instructions", func(t *testing.T) {
		test_string := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
		expected := 48
		result := ExtractDoSubstring(test_string)
		extracted := ExtractMul(result)
		sum := AddExtractedMul(extracted)
		assert.Equal(t, expected, sum)
	})
}
