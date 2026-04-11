package insersection_of_two_arrays

func InsersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, v := range nums1 {
		seen[v] = true
	}
	for _, v := range nums2 {
		if seen[v] {
			result = append(result, v)
			seen[v] = false
		}
	}
	return result
}
