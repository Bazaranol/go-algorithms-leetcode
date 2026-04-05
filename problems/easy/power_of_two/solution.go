package power_of_two

func PowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}

func PowerOfTwoByBit(n int) bool {
	if n > 0 {
		return n&(n-1) == 0
	} else {
		return false
	}
}
