package count_primes

func CountPrimes(n int) int {
	if n < 3 {
		return 0
	}
	arr := make([]bool, n)
	count := 0
	for i := range arr {
		arr[i] = true
	}
	arr[0] = false
	arr[1] = false

	for i := 2; i*i < n; i++ {
		if arr[i] {
			for j := i * i; j < n; j += i {
				arr[j] = false
			}
		}
	}

	for _, value := range arr {
		if value {
			count++
		}
	}

	return count
}
