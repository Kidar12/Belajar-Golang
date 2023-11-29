package main
import "fmt"

func main(){
	var x float64
	var y int
	fmt.Scan(&x)
	y = 0
	for x >= 1 {
		x = x - 1
		y = y + 1
	}
	fmt.Println(x, y)
}