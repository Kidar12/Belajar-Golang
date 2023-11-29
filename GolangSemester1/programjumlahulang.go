package main
import "fmt"

func main() {
	var i, x int
	x = 0
	for i = 1; i <= 99; i += 2 {
		x = x + i
	}
	fmt.Println(x)
}
