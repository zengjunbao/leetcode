package main

import "fmt"

func main() {
	// a := make([]byte, 3, 6)
	// b := make([]byte, 3, 3)
	// equal := bytes.Equal(a, b)
	// fmt.Println(equal)
	//
	// a2 := make([]int, 3, 6)
	// b2 := make([]int, 3, 3)
	// fmt.Println(reflect.DeepEqual(a2, b2))
	// fmt.Println(cap(a2),cap(b2))

	mp := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	for k, v := range mp {
		// if k == 1 {
			delete(mp, 1)
		// }
		fmt.Print(k, v, " ")
	}
	fmt.Println()

}
