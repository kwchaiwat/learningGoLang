package main

import "fmt"

var count int = 0

func main() {
	fmt.Println("Begin")

	// Explicit initialization
	var tmp1 int = 0
	tmp1 = 10
	var tmp2 string = "string"
	var tmp3 bool = false
	const tmp4 int = 0

	// Implicit initialization
	tmp5 := 0
	tmp6 := "string"
	tmp7 := false
	fmt.Println(tmp1, tmp2, tmp3, tmp4, tmp5, tmp6, tmp7)

	for i := 0; i < 5; i++ {
		run()
	}
	println("")

	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}

}

func run() {
	count++
	fmt.Print(count)
}
