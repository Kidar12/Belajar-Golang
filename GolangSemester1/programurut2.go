package main
import "fmt"

func main() {
	var n, x int
	var r, i bool
	fmt.Scan(&n)
	for x = n; x >= n; x--{
		r = x >= n;
		i = x < n
		fmt.Println(r || i)
	}
}
