package main

import "fmt"

func main() {
	var x, y int
	var adik bool
	fmt.Scan(&x, &y)
	adik = (x == y) || (x == 2*y)
	fmt.Println(adik)
}
