func threeSum(nums []int) [][]int {
	var dynamicList [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target, left, right := -nums[i], i+1, len(nums)-1

		// Handle cases where left or right are at the boundaries
		if left == right {
			if right+1 <= len(nums)-1 {
				right++
			} else if left-1 >= 0 {
				left--
			} else {
				break // No more valid combinations
			}
			continue
		}
        
		for left < right {
			sum := nums[right] + nums[left]
			if sum == target {
				newSlice := []int{nums[i], nums[right], nums[left]}
				dynamicList = append(dynamicList, newSlice)

				// Skip duplicate elements
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return dynamicList
}
