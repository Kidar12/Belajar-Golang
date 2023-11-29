package main
import "fmt"

func main() {
	var i, x, n, m float64
	fmt.Scan(&n, &m)
	for i = n; i <= m; {
		x = x + (4 / (i))
		i++
	}
	fmt.Println("%.2f", x)
}
