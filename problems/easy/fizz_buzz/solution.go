package fizz_buzz

import "strconv"

func FizzBuzz(n int) []string {
	arr := make([]string, n)

	for index := range arr {
		num := index + 1
		if num%3 == 0 && num%5 == 0 {
			arr[index] = "FizzBuzz"
		} else if num%3 == 0 {
			arr[index] = "Fizz"
		} else if num%5 == 0 {
			arr[index] = "Buzz"
		} else {
			arr[index] = strconv.Itoa(num)
		}
	}
	return arr
}
