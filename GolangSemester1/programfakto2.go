package main
import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n)
	y = 1
	for x = n; x >= 2; x--{
		y = y * x
	} 
	fmt.Print(y, "x")
}
