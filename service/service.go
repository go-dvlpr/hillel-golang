package service

type Count struct {
	Number int
}

func NewCount() *Count {
	return &Count{
		Number: 0,
	}
}

func (c *Count) Increment() {
	c.Number++
}

func (c *Count) Decrement() {
	c.Number--
}

func (c *Count) Multiplier() {
	c.Number = c.Number * c.Number
}
