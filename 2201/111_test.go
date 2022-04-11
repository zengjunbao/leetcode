package _201

import (
	"fmt"
	"testing"
)

func Test111(t *testing.T)  {
	type name struct {
		id uint64
	}

	a := name{}
	fmt.Printf("%T %v\n",a.id,a.id)
}
