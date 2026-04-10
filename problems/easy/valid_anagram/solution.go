package valid_anagram

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, ch := range s {
		charCount[ch]++
	}

	for _, ch := range t {
		if charCount[ch] <= 0 {
			return false
		}
		charCount[ch]--
	}
	return true
}
