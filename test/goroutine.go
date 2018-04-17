package main

import (
	"fmt"
	//"time"
)

var msg chan int = make(chan int, 10)

func loop(cou int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("cou = %d\n", cou)
	msg <- 1
}

func Afuntion(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {

	for i := 0; i < 10; i++ {
		go loop(i)
	}

	//go loop()
	//time.Sleep(time.Second)

	for i := range msg {
		fmt.Printf("msg = %d", i)
		if len(msg) <= 0 {
			break
		}
	}

	ch := make(chan int) //无缓冲的channel
	//只是把这两行的代码顺序对调一下
	ch <- 1
	go Afuntion(ch)

	// val := <-msg
	// if val == 1 {
	// 	fmt.Printf("channel 接收到数值，值为%d", val)
	// } else {
	// 	fmt.Printf("len(msg) = %d", len(msg))
	// }
}
