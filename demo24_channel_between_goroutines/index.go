package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	c <- payload
	fmt.Println("step 1")
	c <- payload
	fmt.Println("step 2")
	c <- payload
	fmt.Println("step 3")
}

func main() {
	ch := make(chan int, 5)
	go routine1(ch, 1)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
