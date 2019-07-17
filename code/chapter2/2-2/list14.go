package main

import "fmt"

func main() {
	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //无缓冲数据管道

	go consumer(data, done) //启动消费者
	go producer(data)       //启动生产者

	<-done //阻塞，直到消费者发回结束信号

}

// 消费者
func consumer(n chan int, d chan bool) {

	for v := range n { // 接受数据，直到channel关闭
		fmt.Println(v)
	}

	d <- true // 读取结束，通知main协程可以结束程序

}

// 生成者
func producer(n chan int) {
	for i := 0; i < 10; i++ {
		n <- i // 发送数据
	}

	close(n) // 生产数据结束，关闭数据通道
}
