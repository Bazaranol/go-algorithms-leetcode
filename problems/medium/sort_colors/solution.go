package sort_colors

func SortColors(nums []int) []int {
	slow := 0
	medium := 0
	fast := len(nums) - 1
	for medium <= fast {
		if nums[medium] == 0 {
			nums[slow], nums[medium] = nums[medium], nums[slow]
			medium++
			slow++
		} else if nums[medium] == 2 {
			nums[medium], nums[fast] = nums[fast], nums[medium]
			fast--
		} else {
			medium++
		}
	}
	return nums
}
