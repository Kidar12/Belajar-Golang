package main

import "fmt"

func main() {
	var x, x1, x2, x3, x4, x5 int
	fmt.Scan(&x)
	x1 = x / 10000
	x2 = (x / 1000) % 10
	x3 = (x / 100) % 10
	x4 = (x / 10) % 10
	x5 = x % 10
	fmt.Println(x5, x4, x3, x2, x1)
}
