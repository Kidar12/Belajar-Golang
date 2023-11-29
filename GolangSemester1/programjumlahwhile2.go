package main

import "fmt"

func main() {
	var x int
	for i := x; i < 50; i++ {
		x = x + 1
	}
	fmt.Println(x)
}
