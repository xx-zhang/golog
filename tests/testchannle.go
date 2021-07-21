package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var demo = 1

func productor(channel chan<- string) {
	for {
		// 生产环境这里是传入多个line写入
		rData := 10* rand.Float64()
		if rData > 8 {
			channel <- fmt.Sprintf("1111%v", rData)
			channel <- fmt.Sprintf("2222%v", rData + 10 )
			channel <- fmt.Sprintf("3333%v", rData + 100 )
			demo += 3
		}else if rData > 4 {
			channel <- fmt.Sprintf("4444%v", rData - 1 )
			demo += 1
		}else{
			continue
		}
		if demo == 200 {
			close(channel)
		}
		time.Sleep(time.Second * time.Duration(4))
	}
}

func customer(channel <- chan string) {


	for {
		if demo > 4 {
			fmt.Println("开始执行中心函数---------------")
			demo =  0
			}
		message := <-channel  // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
		var i, _ = strconv.Atoi(message)
		if  i > 20 {
			fmt.Println("close ---chinael ")
			}
		}
}

func heatcheck(){
	// 生产环境要增加解释器和数量
	heartbeat := time.Tick(5 * time.Second)

	for {
		select {
		case <- heartbeat:
			//… do heartbeat stuff
			fmt.Println("验证----5-是否该处理-%s", time.Now().Local().String())
			fmt.Println("开始执行中心函数--------demo>4-------")
		}
	}

}

func TestChannel() {
	//var c chan bool
	//c <- true
	//heatcheck(c)

	channel := make(chan string, 5) // 定义带有5个缓冲区的信道(当然可以是其他数字)
	//heartChannel := make(chan int, 3)
	go productor(channel) // 将 productor 函数交给协程处理, 产生的结果传入信道中
	go heatcheck()
	customer(channel) // 主线程从信道中取数据

}