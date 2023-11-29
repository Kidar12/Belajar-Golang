package main
import "fmt"

func main() {
	var i, n, m, x float64
	fmt.Scan(&n, &m)
	for i = n; i <= m; {
		x = x + (2 / (i))
		i++
	}
	fmt.Println("%.2f", x)
}
