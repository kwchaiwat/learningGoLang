package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg

	changeMessage(&msg, "new message 1")
	fmt.Println(*msgPointer)

	changeMessage(msgPointer, "new message 2")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
