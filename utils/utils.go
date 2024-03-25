package utils

//func init() {
//	fmt.Println("Init from utils")
//}

type Counter interface {
	Increment()
	Decrement()
	Multiplier()
}

func SomethingDo(count Counter) {
	count.Increment()
	count.Increment()
	count.Increment()
	count.Increment()
	count.Increment()
}
