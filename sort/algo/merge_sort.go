package algo

func MergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	left := MergeSort(nums[:len(nums)/2])
	right := MergeSort(nums[len(nums)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	l, r := 0, 0
	lSize, rSize := len(left), len(right)
	res := make([]int, 0, lSize+rSize)

	for l < lSize && r < rSize {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	for l < lSize {
		res = append(res, left[l])
		l++
	}

	for r < rSize {
		res = append(res, right[r])
		r++
	}

	return res
}
