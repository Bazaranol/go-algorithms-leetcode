package power_of_four

func PowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func PowerOfFourByBit(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}
