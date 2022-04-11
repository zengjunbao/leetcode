package _111

import (
	"fmt"
	"testing"
)

var list [][]int

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	list = [][]int{}
	var track []int
	backtrack(nums, track)

	return list
}

func backtrack(nums, track []int) {
	// judge end point
	if len(track) == len(nums) {
		list = append(list, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if eleInArr(nums[i], track) {
			continue
		}
		// make choice
		track = append(track, nums[i])
		// backtrack
		backtrack(nums, track)
		// remove choice
		track = track[:len(track)-1]
	}
}

func eleInArr(ele int, list []int) bool {
	for _, num := range list {
		if ele == num {
			return true
		}
	}
	return false
}

func Test46(t *testing.T) {
	fmt.Println("------start----")
	nums := []int{5,4,6,2}
	value := permute(nums)
	fmt.Println("value:", value)
	fmt.Println("-------end-----")
}
