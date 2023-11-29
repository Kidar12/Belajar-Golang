package main
import "fmt"

func main(){
	var i, n, bil, jum int
	var stop bool = false
	fmt.Scan(&n)
	jum = 0; 
	i=1
	for !stop {
		fmt.Scan(&bil)
		jum = jum + bil 
		i++
		stop = i > n
	}
	fmt.Println(jum)
}