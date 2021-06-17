package main

import (
	"demo14/employee"
)

func main() {
	e := employee.New(
		"chaiwat",
		"keawmukdasawan",
		30,
		20)
	e.LeavesRemaining()
}
