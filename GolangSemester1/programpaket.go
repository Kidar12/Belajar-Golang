package main

import "fmt"

func main() {
	var v, p, l, t, b int
	var kirim bool
	fmt.Scan(&p, &l, &t, &b)
	v = (p * l * t )
	b = 30
	kirim = b<30
	fmt.Println(v, kirim)
}
