package container_with_most_water

func ContainerWithMostWater(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for right > left {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
