package single_number

func SingleNumber(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	for num, cnt := range count {
		if cnt == 1 {
			return num
		}
	}
	return 0
}

func SingleNumberByXOR(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
