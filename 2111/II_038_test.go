package _111

import (
	"fmt"
	"testing"
)

func Test_II_038(t *testing.T) {
	// 示例 1:
	// 输入: temperatures = [73,74,75,71,69,72,76,73]
	// 输出: [1,1,4,2,1,1,0,0]

	// 示例 2:
	// 输入: temperatures = [30,40,50,60]
	// 输出: [1,1,1,0]

	// 示例 3:
	// 输入: temperatures = [30,60,90]
	// 输出: [1,1,0]

	// temperatures := []int{30, 60, 90}
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	value := dailyTemperatures(temperatures)
	fmt.Println("value2:", value)
}
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
