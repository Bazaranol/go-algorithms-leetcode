package two_sum_2

func TwoSum2(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] < target {
			left++
		}
		if numbers[left]+numbers[right] > target {
			right--
		}

	}
	return []int{left + 1, right + 1}
}
