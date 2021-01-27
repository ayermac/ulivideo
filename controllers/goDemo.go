package controllers

import (
	"fmt"
	"time"
)

type GoDemoController struct {
	CommonController
}

// @router /go/demo [*]
func (this *GoDemoController) Demo() {
	for i := 0; i < 10; i++ {
		go cal(i)
	}
	time.Sleep(2 * time.Second)
	this.Ctx.WriteString("demo")
}
func cal(i int) {
	fmt.Printf("i = %d\n", i)
}

// @router /go/channel/demo [*]
func (this *GoDemoController) ChannelDemo() {
	Channel := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		//因为假如是无缓冲的channel，在一个goroutine中实现发送和接受就会死锁
		go cal1(i, Channel)
	}
	for j := 0; j < 10; j++ {
		<-Channel
	}
	close(Channel)
	this.Ctx.WriteString("channelDemo")
}
func cal1(i int, Channel chan bool) {
	fmt.Printf("i = %d\n", i)
	time.Sleep(2 * time.Second)
	Channel <- true
}

//select选择
// @router /go/select/demo [*]
func (this *GoDemoController) SelectDemo() {
	oneChannel := make(chan int, 20)
	twoChannel := make(chan string, 20)
	go cal2(oneChannel)

	twoChannel <- "a"
	twoChannel <- "b"

	go func() {
		for {
			select {
			case a := <-oneChannel:
				fmt.Println("get one = ", a)
			case b := <-twoChannel:
				fmt.Println("get two = ", b)
			default:
				fmt.Println("no message")
			}
		}
	}()
	this.Ctx.WriteString("selectDemo")
}
func cal2(oneChannel chan int) {
	for i := 0; i < 10; i++ {
		oneChannel <- i
		fmt.Println("set one = ", i)
	}
}

//模拟任务池
// @router /go/task/demo [*]
func (this *GoDemoController) TaskDemo() {
	//接受任务
	taskChannel := make(chan int, 20)
	//处理任务
	resChannel := make(chan int, 20)
	//关闭任务
	closeChannel := make(chan bool, 5)

	go func() {
		for i := 0; i < 10; i++ {
			taskChannel <- i
		}
		close(taskChannel)
	}()

	//处理任务
	for i := 0; i < 5; i++ {
		go Task(taskChannel, resChannel, closeChannel)
	}

	go func() {
		//当接收到5个值以后，说明5个任务都执行完成了
		for i := 0; i < 5; i++ {
			<-closeChannel
		}
		close(resChannel)
		close(closeChannel)
	}()

	//for循环channel时，当channel关闭以后会退出循环
	for r := range resChannel {
		fmt.Println("res:", r)
	}
	this.Ctx.WriteString("taskDemo")
}
func Task(taskChannel chan int, resChannel chan int, closeChannel chan bool) {
	for t := range taskChannel {
		resChannel <- t
	}
	closeChannel <- true
}
