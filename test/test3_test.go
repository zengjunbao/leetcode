package test

import (
	"fmt"
	"testing"
)

type st struct {
	id   int64
	name string
}

func Test33(t *testing.T) {
	fmt.Println("----------")
	s := &st{}
	test3(s)
	fmt.Println("----------")
}

func test3(s *st) {
	var data []int64
	if s.id != 0 {
		data = append(data, s.id)
	}
	fmt.Println(data)
	fmt.Println(len(data))
}
