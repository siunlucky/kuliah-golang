package main

import "fmt"

func main() {
	var rainfall string

	fmt.Scan(&rainfall)
	if rainfall == "sangat tinggi" {
		fmt.Println("macet")
	} else {
		fmt.Println("tidak macet")
	}
}
