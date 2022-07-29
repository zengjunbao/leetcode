package main

import "fmt"

// https://shuzang.github.io/2020/golang-deep-learning-8-stack-heap-and-escape-analysis/
// 只要有对变量的引用，变量就会存在，而它存储的位置与语言的语义无关。如果可能，变量会被分配到其函数的栈，
// 但如果编译器无法证明函数返回之后变量是否仍然被引用，就必须在堆上分配该变量，采用垃圾回收机制进行管理，
// 从而避免指针悬空。此外，局部变量如果非常大，也会存在堆上。
// 在编译器中，如果变量具有地址，就作为堆分配的候选，但如果逃逸分析可以确定其生存周期不会超过函数返回，就会分配在栈上。
// 总之，分配在堆还是栈完全由编译器确定。

// $ go tool compile -S main.go
// $ go tool compile -m main.go

// func main() {
// 	var a [1]int
// 	c := a[:]
// 	fmt.Println(c)
// }


// func main() {
// 	a := test()
// 	println(a)
// }
//
// func test() *int {
// 	b := 2
// 	return &b
// }

func main() {
	a := make([]int,1,1)
	println(a)
	fmt.Println(a)

	arr := [2]int{1,2}
	fmt.Println(len(arr))
}
