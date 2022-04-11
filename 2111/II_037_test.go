package _111

import (
	"fmt"
	"testing"
)

func Test_II_037(t *testing.T) {
	// asteroids := []int{5, 10, -5}
	// asteroids := []int{-2,-1,1,2}
	// value := asteroidCollision2(asteroids)
	// fmt.Println("value:",value)

	asteroids2 := []int{2,1,-2}
	value2 := asteroidCollision2(asteroids2)
	fmt.Println("value2:",value2)
}

func asteroidCollision(asteroids []int) []int {
	n := len(asteroids)
	if len(asteroids) < 2 {
		return asteroids
	}
	var res []int
	res = append(res, asteroids[0])

	for i := 1; i < n; i++ {
		if asteroids[i]*res[i-1] < 0 {
			if asteroids[i] < res[i-1] {

			}
		}
	}

	return res
}

func asteroidCollision2(asteroids []int) []int {
	var res []int
	for _, a := range asteroids {
		for a < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			sum := res[len(res)-1] + a
			if sum >= 0 {
				a = 0
			}
			if sum < 0 {
				res = res[:len(res)-1]
			}
		}
		if a != 0 {
			res = append(res, a)
		}
	}
	return res
}
