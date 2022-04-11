package sort

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		index := partSort(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
	return arr
}

func partSort(arr []int, left, right int) int {
	// 记录基准
	index := arr[left]
	for left < right {
		// 找出比基准大的值
		for arr[left] < index && left < right {
			left++
		}
		// 找出右边比基准小的值
		for arr[right] > index && left < right {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left] = index
	return left
}
