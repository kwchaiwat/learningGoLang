package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 || someValue < 2 {
		fmt.Println("someValue > 10 || someValue < 2")
	} else {
		fmt.Println("NOT someValue > 10 || someValue < 2")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println("Ok")
	} else {
		fmt.Println("Not ok")
	}
}

func doSomething() string {
	// doSomething
	return "ok"
}
