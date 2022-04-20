package main

import (
	"fmt"
	"reflect"
)

type IPlay interface {
	Ping()
	SetString(s string)
	GetString() string
}
type Player struct {
	name string
}

func (p Player) Ping() {}

func (p Player) SetString(s string) { p.name = s }

func (p Player) GetString() string { return p.name }

func main() {
	// var a int
	// a = 10
	//
	// typeOfA := reflect.TypeOf(a)
	// valueOfA := reflect.ValueOf(a)
	// fmt.Println(typeOfA)
	// fmt.Println(valueOfA, valueOfA.CanSet())
	//
	// valueOfB := reflect.ValueOf(&a)
	// fmt.Println(valueOfB, valueOfB.CanSet())
	//
	// valueOfB = valueOfB.Elem()
	// fmt.Println(valueOfB, valueOfB.CanSet())

	var p Player

	valueOfP := reflect.ValueOf(p)
	fmt.Println(valueOfP, valueOfP.CanSet())

	fmt.Println(valueOfP)
}
