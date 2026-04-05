package happy_number

func HappyNumber(n int) bool {
	prev := make(map[int]bool)

	if n == 1 {
		return true
	}
	for n != 1 {
		if prev[n] {
			return false
		}
		prev[n] = true

		temp := 0
		for n > 0 {
			temp += ((n % 10) * (n % 10))
			n /= 10
		}

		if temp == 1 {
			return true
		}

		n = temp

	}
	return false

}
