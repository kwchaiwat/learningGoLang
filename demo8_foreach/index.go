package main

import "fmt"

func main() {
	var arr1 []int = []int{1, 2, 3, 4}
	var arr2 = []int{1, 2, 3, 4}
	var arr3 [4]string

	arr3[0], arr3[1], arr3[2], arr3[3] = "android", "android", "android", "android"

	for _, item := range arr1 {
		fmt.Print(item, ",")
	}

	for _, item := range arr2 {
		fmt.Print(item, ",")
	}

	for _, item := range arr3 {
		fmt.Print(item, ",")
	}

}
