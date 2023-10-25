package main

import (
    "fmt"
)

func main() {
    var x int;
	var isKk bool;

    fmt.Scan(&x, &isKk)
	fmt.Println("");

	if x >= 17 && isKk == true  {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}



}
