package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n = 5
	for i := 1; i <= n; i++{
		fmt.Println(i)
	}
}