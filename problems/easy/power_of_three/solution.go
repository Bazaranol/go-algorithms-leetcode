package power_of_three

func PowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func PowerOfThreeByBit(n int) bool {
	if n <= 0 {
		return false
	}
	return 1162261467%n == 0
}
