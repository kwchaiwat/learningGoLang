package main

import (
	"demo13/employee"
)

func main() {
	e := employee.Employee{
		FirstName:   "Chaiwat",
		LastName:    "Kaewmukdasawan",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
