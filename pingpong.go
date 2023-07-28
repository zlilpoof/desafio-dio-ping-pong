package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go imprimir(c)
	go pong(c)

	var teclado string

	fmt.Scanln(&teclado)
}

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
