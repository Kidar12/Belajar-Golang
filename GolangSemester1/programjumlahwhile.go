package main

import "fmt"

func main() {
	var x int
	for i := x; x < 1000; i++ {
		x = x + 1
	}
	fmt.Println(x)
}
