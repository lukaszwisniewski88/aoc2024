package fourth

import (
	"testing"
)

func TestSearchFourth(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	t.Run("Count XMAS", func(t *testing.T) {
		expected := 18
		searcher, err := NewSearcher(input, "XMAS")
		if err != nil {
			t.Fatal(err)
		}
		got := searcher.Search()
		if got != expected {
			t.Error("Expected", expected, "got", got)
		}
	})
	t.Run("Count MAS crossed", func(t *testing.T) {
		expected := 9
		searcher, err := NewSearcher(input, "MAS")
		if err != nil {
			t.Fatal(err)
		}
		got := searcher.CrossSearch()
		if got != expected {
			t.Error("Expected", expected, "got", got)
		}
	})
}
