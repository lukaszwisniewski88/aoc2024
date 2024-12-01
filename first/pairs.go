package first

import (
	"math"
	"sort"
)

type Pairs struct {
	left  []int
	right []int
}

func NewPairs() *Pairs {
	return &Pairs{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
}

func (p *Pairs) AddPair(left, right int) {
	p.left = append(p.left, left)
	sort.Slice(p.left, func(i, j int) bool {
		return p.left[i] < p.left[j]
	})
	p.right = append(p.right, right)
	sort.Slice(p.right, func(i, j int) bool {
		return p.right[i] < p.right[j]
	})
}

func (p *Pairs) GetPair(index int) (int, int) {
	return p.left[index], p.right[index]
}

func (p *Pairs) GetDiff(index int) int {
	return int(math.Abs(float64(p.right[index] - p.left[index])))
}

func (p *Pairs) GetRightOccurenceTimes(number int) int {
	count := 0
	for i := 0; i < len(p.right); i++ {
		if p.right[i] == number {
			count++
		}
	}
	return count
}

func (p *Pairs) GetSimilarityScore() int {
	score := 0
	for i := 0; i < len(p.left); i++ {
		score += p.left[i] * p.GetRightOccurenceTimes(p.left[i])
	}
	return score
}

func (p *Pairs) GetAllDiffs() []int {
	diffs := make([]int, 0)
	for i := 0; i < len(p.left); i++ {
		diffs = append(diffs, p.GetDiff(i))
	}
	return diffs
}

func (p *Pairs) GetSumDiffs() int {
	sum := 0
	for i := 0; i < len(p.left); i++ {
		sum += p.GetDiff(i)
	}
	return sum
}
