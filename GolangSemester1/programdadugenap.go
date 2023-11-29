package main

import "fmt"

func main() {
	var x, y, z int
	var hasil bool
	fmt.Scan(&x, &y, &z)
	hasil = (x%2 == 0) && (y%2 == 0) && (z%2 == 0)
	fmt.Println(hasil)
}
