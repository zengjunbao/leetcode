package main

import "fmt"

type stop interface {
	print()
	setFood(food food)
}

type supermarket struct {
	food food
}

func (s *supermarket) print() {
	fmt.Println("---supermarket---")
	fmt.Println(s.food.getFood())
	fmt.Println("---supermarket---")

}

func (s *supermarket) setFood(food food) {
	s.food = food
}

type store struct {
	food food
}

func (s *store) print() {
	fmt.Println("---store---")
	fmt.Println(s.food.getFood())
	fmt.Println("---store---")
}

func (s *store) setFood(food food) {
	s.food = food
}
