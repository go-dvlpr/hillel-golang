package model

type Car struct {
	Color string
	Size  int64
}

func (c Car) Validate() {
	if c.Size == 0 {
		panic("size is 0")
	}

	c.validateColor()
}

func (c Car) validateColor() {
	if c.Color == "" {
		panic("color is empty")
	}
}

//func sayHello() {
//	fmt.Println("hello")
//}
//
//func SayHello() {
//	fmt.Println("hello")
//}
