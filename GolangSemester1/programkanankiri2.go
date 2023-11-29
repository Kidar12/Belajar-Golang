package main
import "fmt"

func main () {
	var x int 
	var n bool
	n = 0
	for n >= 0 {
		fmt.Println(x)
		n = n % 10
	} 
}