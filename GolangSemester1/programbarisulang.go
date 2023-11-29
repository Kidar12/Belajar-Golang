package main

import "fmt"

func main() {
	var i, n, m int
	fmt.Scan(&n, &m)
	i = 0
	for i != n && i != m {
		fmt.Println(i + 1)
	}
}
