package main

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func Test4(t *testing.T) {
	fmt.Println("-------begin------")
	nums := []int{2, 100, 2, 2} // [-2,-2,1,1,4,1,4,4,-4,-2]
	value := singleNumber(nums)
	fmt.Println(value)
	fmt.Println("--------end-------")
}
