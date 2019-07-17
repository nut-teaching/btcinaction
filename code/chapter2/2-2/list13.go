package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	并发
	整个运行时完全并发化设计。凡你能看到的，几乎都在以goroutine方式运行。
	这是一种比普通协程或线程更加高效的并发设计，能轻松创建和运行成千上万的并发任务
 */

var wg sync.WaitGroup

func work(tid int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d:%d\n", tid, i)
		time.Sleep(time.Second)
	}

	wg.Done()
}

func main() {

	wg.Add(2)

	go work(100)
	go work(200)

	wg.Wait()

}
/*
	通道（channel）与goroutine搭配
 */