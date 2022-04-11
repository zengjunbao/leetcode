package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {



	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)










	// n := 10000000
	// ch := make(chan int, 10000000)
	// var a, b, c, d int
	// for i := 0; i < n; i++ {
	// 	select {
	// 	case ch <- i:
	// 		a++
	// 	case ch <- i:
	// 		b++
	// 	case ch <- i:
	// 		b++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	case ch <- i:
	// 		c++
	// 	}
	// }
	// fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)
}
