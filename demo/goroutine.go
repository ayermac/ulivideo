package main

import (
	"fmt"
	"time"
)

func main()  {
	demo1()
}

func demo1()  {
	ch := make(chan int, 1000)
	go func() {
		for i:= 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a:", a)
		}
	}()

	//close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

func demo2() {
	str := []string{"I", "like", "Golang"}

	for k, v := range str {
		go func(i int, s string) {
			fmt.Println(i, s, k, v)
		}(k, v)
	}

	time.Sleep(1e9)
}

func demo3()  {
	str := []string{"I", "like", "Golang"}

	for k, v := range str {
		str = append(str, "good")
		fmt.Println(k, v)
	}

	time.Sleep(1e9)
}

func demo4()  {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		} ()
	}
}

func demo5()  {
	// 使用 `make(chan val-type)` 创建一个新的通道。
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	// 使用 `channel <-` 语法 _发送(send)_ 一个新的值到通道中。这里
	// 我们在一个新的 Go 协程中发送 `"ping"` 到上面创建的
	// `messages` 通道中。
	go func() {
		messages <- "ping"
		//fmt.Println("ping")
	}()

	// 使用 `<-channel` 语法从通道中 _接收(receives)_ 一个值。这里
	// 将接收我们在上面发送的 `"ping"` 消息并打印出来。
	msg := <-messages
	fmt.Println(msg)
	fmt.Println("msg")
}
