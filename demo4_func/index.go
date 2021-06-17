package main

import "fmt"

func main() {
	fmt.Println("Starting")
	name()
	address("KhonKaen, Thailand")
	phone()

	const a, b = 2, 3
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))
	x, y := swapV2(a, b)
	k, j := swap(a, b)
	fmt.Printf("%d,%d => %d,%d\n", a, b, k, j)
	fmt.Printf("%d,%d => %d,%d\n", a, b, x, y)
}

func name() {
	fmt.Println("Chaiwat Kaewmukdasawan")
}
func address(msg string) {
	fmt.Println(msg)
}
func phone() {
	fmt.Println("088-888-xxxx")
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapV2(a int, b int) (x, y int) {
	y = a
	x = b
	return
}
