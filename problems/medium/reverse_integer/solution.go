package reverse_integer

func ReverseInteger(x int) int {
	const INT_MAX = 2147483647
	result := 0
	negative := 1
	if x < 0 {
		negative = -1
		x *= -1
	}
	for x > 0 {
		if result*10+x%10 > INT_MAX {
			return 0
		}
		result = result*10 + x%10
		x /= 10
	}

	return result * negative
}
