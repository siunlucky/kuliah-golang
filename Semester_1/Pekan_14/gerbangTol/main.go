package main

import (
	"fmt"
);

func main()  {
	var x string;
	var a, b, c int;
	var isLoop bool;

	isLoop = true
	a, b, c = 0, 0, 0;


	for isLoop == true{
		fmt.Scan(&x)
		if x == "A" {
			a++
		} else if x == "B" {
			b++
		} else if x == "C" {
			c++
		} else {
			isLoop = false
		}
	}
	
	fmt.Println("")
	fmt.Println("Tipe A =", a)
	fmt.Println("Tipe B =", b)
	fmt.Println("Tipe C =", c)
}