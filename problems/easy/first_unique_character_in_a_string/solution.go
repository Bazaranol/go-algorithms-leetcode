package first_unique_character_in_a_string

func FirstUniqueCharacterInAString(s string) int {
	counter := make(map[rune]int)

	for _, v := range s {
		counter[v]++
	}

	for i, v := range s {
		if counter[v] == 1 {
			return i
		}
	}
	return -1
}
