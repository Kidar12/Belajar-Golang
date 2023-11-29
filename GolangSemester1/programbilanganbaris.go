package main
import "fmt"

func main(){
	var n, m, x int
	fmt.Scan(&n, &m)
	x = 1
	for n = x; n <= m; n++ {
		n = n + 1
	}
	fmt.Println(n)
}