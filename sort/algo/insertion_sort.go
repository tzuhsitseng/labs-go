package algo

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		num := nums[i]
		for j >= 0 && nums[j] > num {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = num
	}
}
