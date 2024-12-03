package third

import (
	"fmt"
	"io"
	"os"
)

func ReadFileAsString(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ProcessDayThree(path string) (string, error) {
	input := ReadFileAsString(path)
	extracted := ExtractMul(input)
	simple_sum := AddExtractedMul(extracted)
	do_extract := ExtractDoSubstring(input)
	do_extracted := ExtractMul(do_extract)
	do_sum := AddExtractedMul(do_extracted)

	return fmt.Sprintf("Simple sum: %d, Do sum: %d", simple_sum, do_sum), nil
}
