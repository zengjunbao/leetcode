package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func Test_0224(t *testing.T) {
	re := -123.98765432123456789
	// re2,_ := decimal.NewFromFloat(re).Round(8).Float64()
	// re3,_ := decimal.NewFromFloat(re).RoundDown(8).Float64()
	// re4,_ := decimal.NewFromFloat(re).RoundUp(8).Float64()
	re5,_ := decimal.NewFromFloat(re).RoundFloor(8).Float64()
	// re6,_ := decimal.NewFromFloat(re).RoundCeil(8).Float64()
	fmt.Printf("%v %T\n",re,re)
	// fmt.Println(re2)
	// fmt.Println(re3)
	// fmt.Println(re4)
	fmt.Println(re5)
	fmt.Println(decimal.NewFromFloat(re5).String())
	fmt.Println(fmt.Sprintf("%v",re5))
	// fmt.Println(re6)
}
