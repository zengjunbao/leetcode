package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Test22222(t *testing.T) {
	arr := []int{3, 2, 1, 4, 7, 6, 5, 8, 11, 12, 13, 15, 10}
	var value []int
	for _, v := range arr {
		wg.Add(1)
		go func(a int) {
			for i := 0; i < a; i++ {
				time.Sleep(time.Millisecond)
			}
			value = append(value, a)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println(value)
}

func Test_2222(t *testing.T) {
	// system V
	// mem, err := shm.NewSystemVMem(6, 10000)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// // mmap
	// mem, err := shm.NewMMapMem("./test.txt", 10000)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}