package palindrome_number

func PalindromeNumber(x int) bool {
	reverse := 0
	num := x

	for num > 0 {
		reverse = reverse*10 + num%10
		num /= 10
	}

	return x == reverse
}
