package sort

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		HeapSortMax(arr, lastLen)
		if i < length {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}
	return arr
}

// 堆排序
func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉树深度
	for i := depth; i >= 0; i-- {
		topMax := i // 假定最大的位置就在i的位置
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if leftChild <= length-1 && arr[leftChild] > arr[topMax] { // 防止越过界限
			topMax = leftChild
		}
		if rightChild <= length-1 && arr[rightChild] > arr[topMax] { // 防止越过界限
			topMax = rightChild
		}
		if topMax != i {
			arr[i], arr[topMax] = arr[topMax], arr[i]
		}
	}
	return arr
}
