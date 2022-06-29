package main


type food interface {
	getFood() string
	setAmount(amount string)
}

type coconut struct {
	name string
}

func (c *coconut) getFood() string {
	return c.name
}

func (c *coconut) setAmount(amount string) {
	c.name = "coconut:"+amount
}

type apple struct {
	name string
}

func (a *apple) getFood() string {
	return a.name
}

func (a *apple) setAmount(amount string) {
	a.name = "apple:"+amount
}
