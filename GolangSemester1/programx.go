package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Println(i % m)
	}
}
