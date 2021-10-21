package main

import (
	"fmt"
	"testing"
)

func number(n int) (m int) {
	for n != 0 {
		if n%2 == 1 {
			m++
		}
		n /= 2
	}
	return
}

func countBits(n int)  []int {
	value := make([]int,n+1,n+1)
	for i := 0; i <= n; i++ {
		value[i] = number(i)
	}
	return value
}

func Test3(t *testing.T) {
	fmt.Println("-------begin------")
	n := 5
	value := countBits(n)
	fmt.Println(value)
	fmt.Println("--------end-------")
}
