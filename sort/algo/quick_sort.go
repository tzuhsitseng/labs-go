package algo

func QuickSort(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, left, right int) {
	if len(nums) <= 1 {
		return
	}
	if right <= left {
		return
	}

	l, r := left, right
	pilot := (left + right) / 2

	for l <= r {
		for l <= r && nums[pilot] > nums[l] {
			l++
		}
		for l <= r && nums[pilot] < nums[r] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	sort(nums, left, r)
	sort(nums, l, right)
}
