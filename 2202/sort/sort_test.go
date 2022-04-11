package sort

import (
	"fmt"
	"testing"
)

func Test_sort(t *testing.T) {
	arr := []int{0, 3, 4, 1, 8, 5, 6, 9, 11, 23, 24, 15,7}
	// fmt.Println(arr)
	QuickSort(arr)
	fmt.Println(arr)
}
