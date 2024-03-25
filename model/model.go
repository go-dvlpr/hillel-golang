package model

import (
	"fmt"
)

func init() {
	fmt.Println("Init from model")
}

func Something() {
	fmt.Println("something")
}

func SomethingElse() {
	fmt.Println("SomethingElse")
}

type Incrementer interface {
	Increment()
}

func SomethingDo(incrementer Incrementer) {
	incrementer.Increment()
	incrementer.Increment()
	incrementer.Increment()
	incrementer.Increment()
	incrementer.Increment()
}

type Decrementer interface {
	Decrement()
}

func SomethingDoElse(decrementer Decrementer) {
	decrementer.Decrement()
	decrementer.Decrement()
	decrementer.Decrement()
	decrementer.Decrement()
	decrementer.Decrement()
}
