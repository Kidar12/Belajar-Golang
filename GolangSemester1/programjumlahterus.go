package main

import "fmt"

func main() {
	var i, x, y, z, a int
	fmt.Scan(&x, &y, &z, &a)
	for i = x; x > 100; x++ {
		x = x * y * z * a
		i = x * i
	}
	fmt.Println(i)
}
