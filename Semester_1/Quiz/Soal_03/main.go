package main

import "fmt"

func main() {
	var x, y string

	fmt.Scan(&x, &y)
	if (x == "merah" && y == "kuning") || (y == "merah" && x == "kuning") {
		fmt.Println("orange")
	} else if (x == "kuning" && y == "biru") || (y == "kuning" && x == "biru") {
		fmt.Println("hijau")
	} else if (x == "biru" && y == "merah") || (y == "biru" && x == "merah") {
		fmt.Println("ungu")
	}
}
