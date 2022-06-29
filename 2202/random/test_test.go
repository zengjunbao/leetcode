package test

import (
	"crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	rand2 "math/rand"
	"testing"
)

func Test2(t *testing.T) {
	var a, b, c int
	for i := 0; i < 10; i++ {
		// str := "ABBCCCCCCC"
		data, _ := rand.Int(rand.Reader, big.NewInt(10))

		fmt.Println(data, data.BitLen())
		// fmt.Printf("%v\n", str[data.Uint64():data.Uint64()+1])
		// if data.Uint64() == 1 {
		// 	a++
		// } else if data.Uint64() > 7 {
		// 	b++
		// } else {
		// 	c++
		// }
	}
	fmt.Println(a, b, c)
}

func Test3(t *testing.T) {
	var a, b, c int
	for i := 1; i < 1000; i++ {
		rand2 := rand2.Intn(100)
		if rand2 < 10 {
			a++
		} else if rand2 > 80 {
			c++
		} else {
			b++
		}

	}
	fmt.Println(a, b, c)
}

func Test4(t *testing.T) {
	for j := 1; j < 100; j++ {
		data, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			fmt.Println("随机出错")
			return
		}
		n1 := decimal.NewFromInt(0)
		n2 := decimal.NewFromInt(1)
		n3 := decimal.NewFromInt(3)
		n4 := decimal.NewFromInt(10)
		n5 := decimal.NewFromInt(10)
		a := []decimal.Decimal{n1, n2, n3, n4, n5}
		// 随机数
		number := decimal.NewFromInt(data.Int64())
		flag := 0
		for i, v := range a {
			if number.LessThanOrEqual(v) {
				flag = i
				break
			}
		}
		fmt.Printf("%d ", flag)
	}

}

func Test1111(t *testing.T)  {
	v := decimal.NewFromFloat(111.11)
	println(v.IntPart())
}