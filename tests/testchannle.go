package tests


import (
"fmt"
"math/rand"
"time"
)

func productor(channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%v", rand.Float64())
		time.Sleep(time.Second * time.Duration(1))
	}
}

func customer(channel <- chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
	}
}

func TestChannel() {
	channel := make(chan string, 5) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	go productor(channel) // 将 productor 函数交给协程处理, 产生的结果传入信道中
	customer(channel) // 主线程从信道中取数据
}