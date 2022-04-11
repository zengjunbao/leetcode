package main

import "fmt"

func main() {
	fmt.Println("----------start----------")
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("触发了宕机, Info =", info)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	reFuncA(-1)
	fmt.Println("-----------end-----------")
}

func reFuncA(a int)  {
	switch a {
	case 0:
		fmt.Println("a panic")
		panic("panic 0")
	case 1:
		defer func() {
			if info := recover(); info != nil {
				fmt.Println("reFuncA宕机, Info =", info)
				panic("panic 2")
			} else {
				fmt.Println("reFuncA正常退出")
			}
		}()
		panic("panic 1")
	default:
		panic("uint err")
	}


}
