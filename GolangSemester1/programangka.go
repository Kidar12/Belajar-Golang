package main

import "fmt"

func main() {
	var x, n, m int
	fmt.Scan(&n, &m)
	x = 1
	for !(n <= m) {
		x = x + 1
	}
	fmt.Print(x)
}
