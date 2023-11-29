package main
import "fmt"

func main() {
	var kar rune
	fmt.Scan(&kar)
	if kar != 'A' || kar != 'I' || kar != 'U' || kar != 'E' || kar != 'O' || kar != 'a' || kar != 'i' || kar != 'u' || kar != 'e' || kar != 'o'{
		fmt.Println("vokal")
	} else {
		fmt.Println("konsonan")
	}
}