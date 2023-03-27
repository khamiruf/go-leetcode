func twoSum(nums []int, target int) []int {
	// use a map to store past items
	seen := make(map[int]int)
	// finding complements
	for i, n := range nums {
		diff := target - n
		if val, ok := seen[diff]; ok {
			return []int{i, val}
		}
		seen[n] = i
	}
	return []int{}
}