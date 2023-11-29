package main
import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	for x = 5; x >= 1; x--{
		fmt.Print(x, "x")
	}
}