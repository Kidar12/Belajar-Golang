package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x%y == 0 {
		fmt.Println("Ya,", x, "kelipatan", y)
	} else {
		fmt.Println("Tidak,", x, "bukan kelipatan", y )
	}
}