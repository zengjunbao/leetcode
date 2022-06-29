package main

import (
	"fmt"
	"math"
	"testing"
)

func divide(a int, b int) int {
	if a == math.MinInt32  && b == -1{
		return math.MaxInt32
	}
	n,m := 0,0
	if a < 0 {
		a = -a
		m++
	}
	if b <0 {
		b = -b
		m++
	}
	for a >= b {
		v1,v2 :=  b,1
		for a >= v1+v1{
			v1 += v1
			v2 += v2
		}
		a -= v1
		n += v2
	}
	if m == 1 {
		return -n
	}
	return n
}

func Test1(t *testing.T) {
	fmt.Println("-------begin------")
	a,b := 15,2
	v := divide(a, b)
	fmt.Println(v)
	fmt.Println("--------end-------")
}

