package ordering

func OrderArr(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		merged[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		merged[i] = right[j]
		i++
	}
	return merged
}
