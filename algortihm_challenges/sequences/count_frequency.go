package sequences

import "strings"

func countFrequency(str string) map[rune]int {
	input := strings.ToLower(str)
	charOccurrences := make(map[rune]int)
	for _, char := range input {
		charOccurrences[char] += 1
	}
	return charOccurrences
}
