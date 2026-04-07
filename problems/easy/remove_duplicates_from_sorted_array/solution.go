package remove_duplicates_from_sorted_array

func RemoveDuplicatesFromSortedArray(nums []int) int {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
