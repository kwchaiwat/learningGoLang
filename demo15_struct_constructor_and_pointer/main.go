package main

import (
	"demo15/employee"
)

func main() {
	e := employee.Inital(
		"chaiwat",
		"keawmukdasawan",
		30,
		20)
	e.LeavesRemaining()

	e = employee.Inital(
		"Lek",
		"Adof",
		30,
		20)
	e.LeavesRemaining()

}
