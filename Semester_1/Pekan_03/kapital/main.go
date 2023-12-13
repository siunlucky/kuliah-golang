package main

import "fmt";

func main() {
	var x string
	var hasil bool

	fmt.Scan(&x)

	hasil = x >= "A" && x <= "Z"
	fmt.Println(hasil)
}