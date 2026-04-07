package maximum_subarray

func MaximumSubarray(nums []int) int {
	localMax := nums[0]
	globalMax := nums[0]

	for i := 1; i < len(nums); i++ {
		localMax = max(localMax+nums[i], nums[i])
		if localMax > globalMax {
			globalMax = max(localMax, globalMax)
		}
	}

	return globalMax
}
