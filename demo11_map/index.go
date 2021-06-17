package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["two"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors["red"])

	var courses = make(map[string]map[string]int)
	// courses["Android"] = map[string]int{"price": 3200, "code": 3213}
	// courses["iOS"] = map[string]int{"price": 4200, "code": 3432}

	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 3200
	courses["Android"]["code"] = 3213

	courses["iOS"] = make(map[string]int)
	courses["iOS"]["price"] = 4200
	courses["iOS"]["code"] = 3432

	fmt.Println("", courses["iOS"]["price"])
}
