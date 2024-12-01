package first

import (
	"testing"
)

func TestPairs(t *testing.T) {
	t.Run("AddPair", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(1, 2)
		if p.left[0] != 1 || p.right[0] != 2 {
			t.Errorf("Expected 1, 2 but got %d, %d", p.left[0], p.right[0])
		}
	})
	t.Run("Added pairs should be sorted", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(2, 1)
		p.AddPair(1, 5)
		p.AddPair(3, 4)
		p.AddPair(0, 0)
		for i := 0; i < len(p.left)-1; i++ {
			if p.left[i] > p.left[i+1] || p.right[i] > p.right[i+1] {
				t.Errorf("Expected sorted left but got %v", p.left)
			}
		}
	})
	t.Run("Count the difference between pairs", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(1, 2)
		p.AddPair(2, 1)
		p.AddPair(3, 4)
		p.AddPair(4, 3)
		p.AddPair(46, 12)

		if p.GetDiff(0) != 0 || p.GetDiff(1) != 0 || p.GetDiff(2) != 0 || p.GetDiff(3) != 0 || p.GetDiff(4) != 34 {
			t.Errorf("Expected 1 but got %d", p.GetDiff(0))
		}
	})
	t.Run("Get Sum differences", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(3, 4)
		p.AddPair(4, 3)
		p.AddPair(2, 5)
		p.AddPair(1, 3)
		p.AddPair(3, 9)
		p.AddPair(3, 3)
		if p.GetSumDiffs() != 11 {
			t.Errorf("Expected 11 but got %d", p.GetSumDiffs())
		}
	})
	t.Run("Get how many times a number occurs in the right side", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(3, 4)
		p.AddPair(4, 3)
		p.AddPair(2, 5)
		p.AddPair(1, 3)
		p.AddPair(3, 9)
		p.AddPair(3, 3)
		if p.GetRightOccurenceTimes(3) != 3 {
			t.Errorf("Expected 3 but got %d", p.GetRightOccurenceTimes(3))
		}
	})
	t.Run("Get similarity score", func(t *testing.T) {
		p := NewPairs()
		p.AddPair(3, 4)
		p.AddPair(4, 3)
		p.AddPair(2, 5)
		p.AddPair(1, 3)
		p.AddPair(3, 9)
		p.AddPair(3, 3)
		if p.GetSimilarityScore() != 31 {
			t.Errorf("Expected 38 but got %d", p.GetSimilarityScore())
		}
	})
}
