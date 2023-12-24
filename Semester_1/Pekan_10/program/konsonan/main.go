package main

import (
    "fmt"
)

func main() {
    var x string;

    fmt.Scan(&x)
	fmt.Println("");

	if x == "a" || x == "A" || x == "e" || x == "E" || x == "i" || x == "I" || x == "o" || x == "O" || x == "u" || x == "U" {
		fmt.Println("bukan konsonan")
	} else {
		fmt.Println("konsonan")
	}

}
