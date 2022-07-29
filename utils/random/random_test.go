package random

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var GiftEx = []Gift{
	{
		Id:     1,
		Weight: 0.01,
		Type:   "S",
	}, {
		Id:     2,
		Weight: 0.09,
		Type:   "A",
	}, {
		Id:     3,
		Weight: 0.2,
		Type:   "B",
	}, {
		Id:     4,
		Weight: 0.3,
		Type:   "C",
	}, {
		Id:     5,
		Weight: 0.4,
		Type:   "D",
	},
}

func TestRandom(t *testing.T) {
	var s, a, b, c, d int
	var num = 10000000
	for i := 0; i < num; i++ {
		random, err := Random(GiftEx)
		assert.NoError(t, err)
		switch random.Type {
		case "S":
			s++
		case "A":
			a++
		case "B":
			b++
		case "C":
			c++
		case "D":
			d++
		}
	}
	fmt.Println("S:", s, float64(s*100)/float64(num))
	fmt.Println("A:", a, float64(a*100)/float64(num))
	fmt.Println("B:", b, float64(b*100)/float64(num))
	fmt.Println("C:", c, float64(c*100)/float64(num))
	fmt.Println("D:", d, float64(d*100)/float64(num))

}

func TestRandomLottery(t *testing.T) {
	mp1 := make(map[int]int)
	for i := 0; i < 10000; i++ {
		lottery, _ := Lottery2()
		mp1[lottery.ID] ++
	}
	fmt.Println(mp1)

	mp2 := make(map[int]int)
	for i := 0; i < 10000; i++ {
		lottery, _ := Lottery2()
		mp2[lottery.ID] ++
	}
	fmt.Println(mp2)
}
