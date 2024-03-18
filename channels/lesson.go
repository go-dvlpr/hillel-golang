package channels

import (
	"fmt"
	"sync"
	"time"
)

func Start() {
	//ch := make(chan string)
	//chInt := make(chan int)

	//go FirstPublisher(ch)
	//go SecondPublisher(chInt)

	//go FirstSubscriber(ch)
	//go SecondSubscriber(ch)
	//go ThirdSubscriber(ch, chInt)
	m = make(map[int]int)

	go WriteToMap()
	go WriteToMap()
	go WriteToMap()
	go WriteToMap()
	go WriteToMap()

	time.Sleep(time.Second * 3)

	fmt.Println(m)
}

var m = map[int]int{}
var mut = sync.RWMutex{}

func WriteToMap() {
	for i := 0; i < 1000; i++ {
		mut.Lock()
		m[i] = i * 2
		mut.Unlock()
	}
}

func FirstPublisher(ch chan<- string) {
	for i := 0; i < 10; i++ {
		//fmt.Println("prepearing to send ", i)
		ch <- fmt.Sprintf("something %d", i)
		time.Sleep(time.Millisecond * 500)
	}

	close(ch)
}

func SecondPublisher(chInt chan<- int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("prepearing to send ", i)
		chInt <- i
		time.Sleep(time.Millisecond * 700)
	}

	close(chInt)
}

func FirstSubscriber(ch chan string) {
	count := 0
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("FirstSubscriber: channel was closed")
			return
		}

		fmt.Println("FirstSubscriber recived: ", data, " ---> isOpen: ", ok)
		count++
		if count == 10 {
			break
		}
	}
}

func SecondSubscriber(ch <-chan string) {
	for data := range ch {
		fmt.Println("SecondSubscriber recived: ", data)
	}

	fmt.Println("SecondSubscriber: channel was closed")
}

func ThirdSubscriber(ch <-chan string, chInt <-chan int) {
	tickerChannel := time.NewTicker(time.Second).C

	for {
		select {
		case data, ok := <-ch: //12:30:50.100000 (random)
			if !ok {
				fmt.Println("ThirdSubscriber: string channel was closed")
				return
			}

			fmt.Println("ThirdSubscriber: string recieved: ", data, " --> isOpen: ", ok)

		case data, ok := <-chInt: //12:30:50.100000 (random)
			if !ok {
				fmt.Println("ThirdSubscriber: int channel was closed")
				return
			}

			fmt.Println("ThirdSubscriber: int recieved: ", data, " --> isOpen: ", ok)

		case c := <-tickerChannel: //12:30:50.100000 (random)
			fmt.Println("tick, ", c)
		}
	}
}
