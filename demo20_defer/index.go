package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer printFinished(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}
}

func printFinished(i int) {
	fmt.Println("Finish: ", i)
}
