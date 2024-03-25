package utils

import (
	"fmt"
	"testing"
)

type MockedCounter struct {
	Number int
}

func (c *MockedCounter) Increment() {
	c.Number++
	c.Number++
	c.Number++
	c.Number++
}

func (c *MockedCounter) Decrement() {
	c.Number--
	c.Number--
	c.Number--
}

func (c *MockedCounter) Multiplier() {
	c.Number = c.Number * c.Number * c.Number
}

func TestSomethingDo(t *testing.T) {

	mc := &MockedCounter{}

	SomethingDo(mc)

	fmt.Println(mc)
}
