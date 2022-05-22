package algo

func BubbleSort(nums []int) {
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				sorted = false
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
