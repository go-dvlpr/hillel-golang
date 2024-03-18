package model

import "fmt"

type Voicer interface {
	Voice()
}

var _ Voicer = (*Dog)(nil)

type Dog struct{}

func (d Dog) Voice() {
	fmt.Println("woof woof")
}

var _ Voicer = (*Cat)(nil)

type Cat struct{}

func (d Cat) Voice() {
	fmt.Println("meow meow")
}

var _ Voicer = (*Horse)(nil)

type Horse struct{}

func (d Horse) Voice() {
	fmt.Println("something")
}
