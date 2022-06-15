package _struct

import (
	"fmt"
	"testing"
)

func TestCopyStruct(t *testing.T) {
	a := OrderA{
		Id:      1,
		Name:    "test1",
		Address: "adadada",
		Phone:   "12312341234",
	}
	b := OrderB{}
	err := CopyStruct(&a, &b)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("b:",b.Id)
}
