package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	quit := make(chan int)
	go looper(quit)
	for scanner.Scan() {
		if scanner.Text() == "stop" {
			quit <- 0
			fmt.Println("hello")
			return
		}
	}
}

func looper(quit chan int) {
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Println("asd")
			time.Sleep(1000 * time.Millisecond)
		}

	}
}
