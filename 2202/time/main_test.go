package main

import (
	"fmt"
	"testing"
)

func Test111(t *testing.T) {
	str := "s2"
	switch str {
	case "s1":
		fallthrough
	case "s2":
		fallthrough
	case "s3":
		fallthrough
	case "s4":
		fmt.Println("s4:", str)
	}
}
