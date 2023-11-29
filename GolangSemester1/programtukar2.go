package main

import "fmt"

func main() {
	var x, x1, x2, x3, x4 int
	fmt.Scan(&x)
	x1 = x / 1000								
	x2 = (x / 100) % 10						
	x3 = (x / 10) % 10
	x4 = x % 10
	fmt.Println(x4, x3, x2, x1)
}
