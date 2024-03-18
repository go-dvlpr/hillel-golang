package main

import (
	"fmt"
	"hillel_new/channels"
	"hillel_new/model"
	"sync"
	"time"
)

// main -> counter 1
// main -> counter 2

// main -> runner -> counter 1
// main -> runner -> counter 2

func main() {
	channels.Start()

	h := model.Horse{}
	fmt.Println(h)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go runner(wg)

	wg.Wait()

	fmt.Println("finished")

}

func runner(wg *sync.WaitGroup) {
	go Counter("first", wg)
	go Counter("second", wg)
	fmt.Println("runner finished")
}

func Counter(name string, wg *sync.WaitGroup) {
	count := 0
	for {
		if count == 10 {
			break
		}
		count++

		//fmt.Println(name, " --> ", count)

		time.Sleep(time.Second)
	}

	wg.Done()

}
