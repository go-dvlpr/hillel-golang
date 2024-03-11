package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	datamodel "hillel_new/model"
)

//func main() {
//TestFunction("2", "red", "green")

//books := make(map[string]int)
//books["book1"] = 100
//books["book2"] = 200
//books["book3"] = 300
//
//read(books)

//defer func() {
//	//r := recover()
//	//if r != nil {
//	//	fmt.Println("something 1")
//	//}
//
//	if r := recover(); r != nil {
//		fmt.Println("something")
//	}
//}()

//a := []int{1, 2, 3}
//fmt.Println(a[2])

//bird := "sinichka"
//if bird == "sinichka" {
//	panic("bird is not sinichka")
//}
//
//fmt.Println("after panic")

//pet := Pet{
//	name:  "dmytro",
//	color: "red",
//}
//
//pet2 := NewPet("alex", "white")
//
//fmt.Println(pet)
//fmt.Println(pet2)

//pet := new(Pet)
//fmt.Println(pet)
//var pet2 *Pet
//fmt.Println(pet2)

//	car := model.LegkovaAuto{Color: "blue"}
//	fmt.Println(car)
//	fmt.Println(model.Variable1)
//	model.Hello()
//
//}

//type Pet struct {
//	name  string
//	color string
//}

//func NewPet(name, color string) *Pet {
//	return &Pet{
//		name:  name,
//		color: color,
//	}
//}

//func read(books map[string]int) {
//	book1, ok := books["book3"]
//
//	fmt.Println(ok)
//	fmt.Println(book1)
//}

//func TestFunction(a string, color ...string) {
//	for i, v := range color {
//		fmt.Println(i, v)
//	}
//}

func main() {
	car := datamodel.Car{
		Color: "yellow",
		Size:  1,
	}

	fmt.Println(car)

	car.Validate()

	var sl = []int{1, 2, 3, 4}

	element, err := GetElementFromSlice(sl, 5)
	if err != nil {
		logrus.Error("failed to get element from slice - ", err)
		return
	}

	element++ //
}

func GetElementFromSlice(sl []int, index int) (int, error) {
	if len(sl)-1 < index {
		return 0, errors.New("index out of range")
	}

	element := sl[index]
	return element, nil
}
