package subarray_sum_equals_k

func SubarraySumEqualsK(nums []int, k int) int {
	curSum, count := 0, 0
	sumMap := make(map[int]int)
	sumMap[0] = 1

	for _, num := range nums {
		curSum += num

		if val, exists := sumMap[curSum-k]; exists {
			count += val
		}

		sumMap[curSum]++
	}
	return count
}
