package main
import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if y-x == z-y {
		fmt.Println("Barisan aritmatika")
	} else {
		if y/x == z/y {
			fmt.Println("Barisan geometri")
		} else {
			fmt.Println("Bukan aritmatika atau geometri")
		}
	}
}
