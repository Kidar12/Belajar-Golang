package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x >= 1000 {
        if y >= 4000 {
            fmt.Println("sudah dapat dimonetisasi")
        } else {
            fmt.Println("belum dapat dimonetisasi")
        }
    } else {
        fmt.Println("belum dapat dimonetisasi")
    }
}
