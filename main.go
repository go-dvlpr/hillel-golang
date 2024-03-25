package main

import (
	"fmt"
	"hillel_new/model"
	"hillel_new/service"
	"hillel_new/utils"
	"sync"
)

var m map[int]int
var mSync sync.Map

func main() {
	counter := service.NewCount()

	utils.SomethingDo(counter)
	model.SomethingDo(counter)
	model.SomethingDoElse(counter)

	fmt.Println(counter)
}

//func main() {
//	fmt.Println("hello")
//
//	//model.Something()
//	//model.SomethingElse()
//
//	once := sync.Once{}
//	once.Do(model.Something)
//	once.Do(model.Something)
//
//	f := sync.OnceFunc(model.SomethingElse)
//	f()
//	f()
//
//	m = make(map[int]int)
//
//	for i := 0; i < 100; i++ {
//		go Write(i)
//	}
//
//	time.Sleep(time.Second)
//
//	fmt.Println(mSync.Load(5))
//	mSync.Delete(5)
//	fmt.Println(mSync.Load(5))
//
//	var count int32 = 0
//
//	for i := 0; i < 100; i++ {
//		go func() {
//			atomic.AddInt32(&count, 10)
//		}()
//	}
//
//	time.Sleep(time.Second)
//
//	fmt.Println(count)
//
//	mut := sync.Mutex{}
//
//	count = 0
//	start := time.Now()
//	for i := 0; i < 1000000; i++ {
//		go func() {
//			mut.Lock()
//			count++
//			mut.Unlock()
//		}()
//	}
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//
//	count = 0
//	start = time.Now()
//	for i := 0; i < 1000000; i++ {
//		atomic.AddInt32(&count, 1)
//	}
//	end = time.Now()
//	fmt.Println(end.Sub(start))
//
//}
//
//func Write(i int) {
//	//m[i] = i * 2
//	mSync.Store(i, i*2)
//}

//func main() {
//ctx := context.Background()
//ctx = context.TODO()
//
//ctx = context.WithValue(context.Background(), "something", "here")
////ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
////ctx, cancel := context.WithTimeout(ctx, time.Second*3)
//
//ctx, cancel := context.WithCancel(ctx)
//
//go Something(ctx)
//
//time.Sleep(time.Second * 1)
//
//cancel()
//
//result := SomethingTwo()
//fmt.Println(result)

//}

//func SomethingTwo() (a string) {
//	a = "Hello from SomethingTwo faunction"
//	return
//}
//
//func Something(ctx context.Context) {
//	res := ctx.Value("something")
//	fmt.Println(res)
//
//	//resDone := ctx.Done()
//	//start := time.Now()
//	//end := time.Now()
//
//	//fmt.Println(end.Sub(start))
//
//	for {
//		select {
//		case <-time.NewTicker(time.Second).C:
//			fmt.Println(time.Now())
//		case <-ctx.Done():
//			fmt.Println("done")
//			return
//		}
//	}
//
//}
