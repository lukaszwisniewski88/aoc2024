package second

import "testing"

func TestIsReportSafe(t *testing.T) {
	lines := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := []bool{true, false, false, false, false, true}
	for i, line := range lines {
		if IsReportSafe(line) != expected[i] {
			t.Errorf("Expected %t, got %t", expected[i], IsReportSafe(line))
		}
	}

}
func TestCountSafeReports(t *testing.T) {
	lines := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := 2
	if CountSafeReports(lines) != expected {
		t.Errorf("Expected %d, got %d", expected, CountSafeReports(lines))
	}
}
func TestCountSafeByRemoving(t *testing.T) {
	lines := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := 4
	if CountLessSafeReports(lines) != expected {
		t.Errorf("Expected %d, got %d", expected, CountLessSafeReports(lines))
	}
}
