package main

import "fmt"

func main() {
	forLoop()
	whileLoop()
	foreachLoop()
}

func forLoop() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Index %d\n", i)
	}
}

func whileLoop() {
	i := 0
	for i < 4 {
		i++
		fmt.Printf("Index %d\n", i)
	}
}

func foreachLoop() {
	courses := []string{"Android", "iOS", ".Net", "VueJs", "React", "PHP"}
	for index, course := range courses {
		if index == 3 {
			break
		}
		fmt.Printf("index: %d, Course: %s\n", index+1, course)
	}
}
