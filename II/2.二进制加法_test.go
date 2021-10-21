package main

import (
	"fmt"
	"testing"
)

func addBinary(a string, b string) string {
	n1, n2, m := len(a), len(b), 0
	n0 := maxNum(n1, n2)
	str := ""
	for i := 0; i < n0; i++ {
		if i < n1 {
			if a[n1-i-1] == '1' {
				m++
			}
		}
		if i < n2 {
			if b[n2-i-1] == '1' {
				m++
			}
		}
		if m%2 == 1 {
			str = "1" + str
			m -= 1
		} else {
			str = "0" + str
		}
		m /= 2
	}
	if m > 0 {
		str = "1" + str
	}
	return str
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test2(t *testing.T) {
	fmt.Println("-------begin------")
	a, b := "11", "10"
	str := addBinary(a, b)
	fmt.Println(str)
	fmt.Println("--------end-------")
}