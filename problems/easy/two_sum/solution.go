package twosum

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		prevIndex, ok := seen[complement]
		if ok {
			return []int{prevIndex, index}
		}
		seen[value] = index
	}
	return nil
}
