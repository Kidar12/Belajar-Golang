package main
import "fmt"

func main () {
	var input, bil1, bil2, bil3, bil4, bil5 int
	fmt.Scan(&input)

	bil1 = input / 10000
	bil2 = input / 1000 % 10
	bil3 = input / 100 % 10
	bil4 = input / 10 % 10
	bil5 = input % 10

	fmt.Printf("%d %d %d %d %d\n", bil5, bil4, bil3, bil2, bil1)
}