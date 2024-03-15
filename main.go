package main

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"hillel_new/errs"
	"log"
	"math/big"
	"net/http"

	//"math/rand/v2"
	//"math/rand/v2"
	"crypto/rand"
)

type Sleeper interface {
	Sleep()
}

type Human struct{}

func (h Human) Sleep() {
	fmt.Println("Human is sleeping right now")
}

type Pet struct{}

func (h Pet) Sleep() {
	fmt.Println("Pet is sleeping right now")
}

func GoToBed(sleeper Sleeper) {
	sleeper.Sleep()
}

func GoToBedHuman(human Human) {
	human.Sleep()
}

func GoToBedPet(pet Pet) {
	pet.Sleep()
}

func H(w http.Request, r *http.Response) {

}

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Info("congratulations, main was started")

	animal, err := NewAnimal("")
	if err != nil {
		if errors.Is(err, errs.ErrEmptyName) {
			//logger.Error("animal has empty name")
			logger.Fatal("finish")
			//return
			//logger.Warn("animal has empty name")
		}

		log.Println("failed to create animal,", err)
		//return
	}

	// debug
	// info
	// warning
	// error
	// fatal

	//logger.Debug(fmt.Sprintf("animal has name %s", animal.Name))

	log.Println("animal created successfuly, ", animal)

	PrintSomething(1)
	PrintSomething("test")
	PrintSomething(false)
	PrintSomething(animal)

	PrintSomethingForRunner(animal)

	human := Human{}
	pet := Pet{}

	GoToBedHuman(human)
	GoToBedPet(pet)

	GoToBed(human)
	GoToBed(pet)

	PrintInterfaceType(animal)
}

func PrintInterfaceType(something interface{}) {
	// if string -> this is string
	// if int -> this is int
	// if Animal -> this is Animal

	strSomething, ok := something.(string)
	if ok {
		fmt.Println("this is string", strSomething)
	}

	intSomething, ok := something.(int)
	if ok {
		fmt.Println("this is int", intSomething)
	}

	boolSomething, ok := something.(bool)
	if ok {
		fmt.Println("this is bool", boolSomething)
	}

	switch a := something.(type) {
	case string:
		fmt.Println("this is string again", a)
	case int:
		fmt.Println("this is int again", a)
	case bool:
		fmt.Println("this is bool again", a)
	case *Animal:
		fmt.Println("oooh shit this is animal", a)
	}
}

type Runner interface {
	RunFast() error
}

func PrintSomething(arg interface{}) {
	fmt.Println(arg)
}

func PrintSomethingForRunner(arg Runner) {
	fmt.Println(arg)
}

type Animal struct {
	Name string
}

func NewAnimal(name string) (*Animal, error) {
	if name == "" {
		//err := fmt.Errorf("name is empty %d", 12223546)
		//
		//errorText := fmt.Sprintf("name is empty %d", 12223546)
		//err = errors.New(errorText)
		customErr := errs.ErrEmptyName

		return nil, customErr
	}

	return &Animal{Name: name}, nil
}

func (a Animal) RunFast() error {
	max := big.NewInt(1000)
	randomNumber, _ := rand.Int(rand.Reader, max)
	res := randomNumber.Int64() + 1000

	if res > 1100 && res < 1800 {
		return errors.New("something went wrong")
	}

	fmt.Println("RunFast called")

	return nil
}
