package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c)
	if a >= b && a >= c {
		d = a
	} else if b >= a && b >= c {
		d = b
	} else {
		d = c
	}
	fmt.Println(d)
}
