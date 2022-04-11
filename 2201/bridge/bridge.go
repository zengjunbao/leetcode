package main

func main() {
	coconut := &coconut{}
	apple := &apple{}

	supermarket := &supermarket{}
	supermarket.setFood(coconut)
	supermarket.food.setAmount("￥22")
	supermarket.print()

	supermarket.setFood(apple)
	supermarket.food.setAmount("￥15")
	supermarket.print()

	store := &store{}
	store.setFood(coconut)
	store.food.setAmount("$22")
	store.print()

	store.setFood(apple)
	store.food.setAmount("$15")
	store.print()
}
