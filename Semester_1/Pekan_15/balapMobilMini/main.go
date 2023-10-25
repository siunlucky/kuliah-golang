package main

import "fmt"

func main() {
    var x, first string;
	var a, b int;

	a, b = 0, 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&x)
		if x == "A" {
			a++
		} else if x == "B" {
			b++
		}

		if a == 5 && first == "" {
			first = "A"
		} else if b == 5 && first == "" {
			first = "B"
		} else if b > 5 || a > 5 {
			first = "tidak valid"
		}
	}
	fmt.Println("")
	fmt.Println(first)

}
