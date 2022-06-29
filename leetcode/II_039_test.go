package _111

import (
	"fmt"
	"testing"
)

func Test_II_039(t *testing.T) {
	// 输入：heights = [2,1,5,6,2,3]
	// 输出：10
	// 输入： heights = [2,4]
	// 输出： 4
	// heights := []int{2,1,5,6,2,3}
	heights := []int{2, 4}
	value := largestRectangleArea1(heights)
	fmt.Println("value:", value)
}

func largestRectangleArea1(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	var monoStack []int
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}

	fmt.Println("left:",left)
	fmt.Println("right:",right)

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	var monoStack []int
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
