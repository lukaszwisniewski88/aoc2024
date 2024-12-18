package fourth

import (
	"fmt"
	"lukaszwisniewski88/aoc2024/first"
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

type Vec2 struct {
	x, y int
}

type Searcher struct {
	matrix   []string
	sequence string
}

func NewSearcher(input []string, sequence string) (*Searcher, error) {
	if sequence == "" || len(input) == 0 {
		return nil, fmt.Errorf("invalid input")
	}
	return &Searcher{
		// matrix[y][x]
		matrix:   input,
		sequence: strings.TrimSpace(sequence),
	}, nil
}

func (s *Searcher) hasSequence(x, y int, v Vec2) bool {
	// if first letter is not the same then return false
	// vector to sequence of positions
	searching_sequence := s.sequence

	for i := 0; i < len(searching_sequence); i++ {
		vec_x := v.x / (len(searching_sequence) - 1)
		vec_y := v.y / (len(searching_sequence) - 1)
		letter, err := s.GetChar(x+(i*vec_x), y+(i*vec_y))
		if err != nil {
			fmt.Println("Error getting char", err)
			return false
		}
		if letter != rune(searching_sequence[i]) {
			return false
		}
	}
	return true
}

func (s *Searcher) hasCrossSequence(x, y int) bool {
	// if first letter is not the same then return false
	// vector to sequence of positions
	center_letter := "A"
	// only if the letter is the center letter we can continue
	actual_letter, err := s.GetChar(x, y)
	if err != nil {
		return false
	}
	if actual_letter != rune(center_letter[0]) {
		return false
	}
	all_vectors := []Vec2{
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}
	// filter out impossible vectors
	possible_vectors := make([]Vec2, 0)
	for _, vec := range all_vectors {
		if x+vec.x >= 0 && x+vec.x < len(s.matrix[y]) && y+vec.y >= 0 && y+vec.y < len(s.matrix) {
			possible_vectors = append(possible_vectors, vec)
		}
	}
	if len(possible_vectors) != len(all_vectors) {
		return false
	}
	// check if the sequence is correct
	for i := 0; i < 2; i++ {
		actual_vector := possible_vectors[i*2]
		opposite_vector := Vec2{actual_vector.x * -1, actual_vector.y * -1}

		actual_vector_char, err := s.GetChar(x+actual_vector.x, y+actual_vector.y)
		if err != nil {
			return false
		}
		opposite_vector_char, err := s.GetChar(x+opposite_vector.x, y+opposite_vector.y)
		if err != nil {
			return false
		}
		if actual_vector_char == opposite_vector_char {
			return false
		}
		if !(actual_vector_char == rune("M"[0]) || actual_vector_char == rune("S"[0])) || !(opposite_vector_char == rune("M"[0]) || opposite_vector_char == rune("S"[0])) {
			return false
		}

	}
	return true
}

func (s *Searcher) getPossibleVectors(x, y, max_x, max_y int) []Vec2 {
	// get the length of the sequence
	sequence_len := len(s.sequence) - 1
	// get the length of the alternate sequence
	vectors := make([]Vec2, 0)
	// add the possible vectors
	if x+sequence_len < max_x {
		vectors = append(vectors, Vec2{sequence_len, 0})
	}
	if x-sequence_len >= 0 {
		vectors = append(vectors, Vec2{-sequence_len, 0})
	}
	if y+sequence_len < max_y {
		vectors = append(vectors, Vec2{0, sequence_len})
	}
	if y-sequence_len >= 0 {
		vectors = append(vectors, Vec2{0, -sequence_len})
	}
	if x+sequence_len < max_x && y+sequence_len < max_y {
		vectors = append(vectors, Vec2{sequence_len, sequence_len})
	}
	if x-sequence_len >= 0 && y-sequence_len >= 0 {
		vectors = append(vectors, Vec2{-sequence_len, -sequence_len})
	}
	if x+sequence_len < max_x && y-sequence_len >= 0 {
		vectors = append(vectors, Vec2{sequence_len, -sequence_len})
	}
	if x-sequence_len >= 0 && y+sequence_len < max_y {
		vectors = append(vectors, Vec2{-sequence_len, sequence_len})
	}

	return vectors
}

func (s *Searcher) GetChar(x, y int) (rune, error) {
	// check if x and y are within the bounds of the matrix
	if x < 0 || x > len(s.matrix[y]) || y < 0 || y > len(s.matrix) {
		return 0, fmt.Errorf("out of bounds")
	}
	return rune(s.matrix[y][x]), nil
}

func (s *Searcher) Search() (count int) {
	for y, row := range s.matrix {
		for x := range row {
			// get possible vectors in this position
			vectors := s.getPossibleVectors(x, y, len(row), len(s.matrix))
			for _, vector := range vectors {
				if s.hasSequence(x, y, vector) {
					count++
				}
			}
		}
	}
	return count
}

func (s *Searcher) CrossSearch() (count int) {
	for y, row := range s.matrix {
		for x := range row {
			// we will have allways the same cross vectors
			if s.hasCrossSequence(x, y) {
				count++

			}
		}
	}
	return
}

func ProcessDayFour(input_path string) (string, error) {
	scanner, err := first.OpenFile(input_path)
	if err != nil {
		return "", err
	}
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	searcher, err := NewSearcher(input, "XMAS")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d \n %d", searcher.Search(), searcher.CrossSearch()), nil
}
