package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var kota1, kota2 string
	fmt.Scan(&kota1, &kota2)
	// Tentukan apakah bertetanggaan
	if kota1 == "Bogor" && kota2 == "Sukabumi" {
		fmt.Println(bertetangga)
	} else if kota1 == "Sukabumi" && kota2 == "Bogor" {
		fmt.Println(bertetangga)
	} else if kota1 == "Bogor" && kota2 == "Cianjur" {
		fmt.Println(bertetangga)
	} else if kota1 == "Cianjur" && kota2 == "Bogor" {
		fmt.Println(bertetangga)
	} else if kota1 == "Sukabumi" && kota2 == "Cianjur" {
		fmt.Println(bertetangga)
	} else if kota1 == "Cianjur" && kota2 == "Sukabumi" {
		fmt.Println(bertetangga)
	} else if kota1 == "Cianjur" && kota2 == "Tasikmalaya" {
		fmt.Println(tidak bertetangga)
	} else if kota1 == "Tasikmalaya" && kota2 == "Cianjur" {
		fmt.Println(tidak bertetangga)
	} else if kota1 == "Tasikmalaya" && kota2 == "Bandung" {
		fmt.Println(bertetangga)
	} else if kota1 == "Bandung" && kota2 == "Tasikmalaya" {
		fmt.Println(bertetangga)
	} else {
		fmt.Println(bertetangga)
	}

	// Tampilkan output
	fmt.Println(bertetangga)
}
