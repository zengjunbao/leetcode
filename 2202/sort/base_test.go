package sort

import (
	"fmt"
	"testing"
)

func Test_Base(t *testing.T) {
	var arr [3][]int
	myarr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3}
	for i := 0; i < len(myarr); i++ {
		arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
	}
	fmt.Println(arr)
}

func Test_Base(t *testing.T) {
	var arr [3][]int
	myarr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3}
	for i := 0; i < len(myarr); i++ {
		arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
	}
	fmt.Println(arr)
}

