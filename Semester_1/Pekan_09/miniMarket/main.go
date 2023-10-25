package main

import (
    "fmt"
)

func main() {
    var x int
    var isCard, isDiscount, isCashback bool

    fmt.Scan(&x, &isCard)
	fmt.Println("");

	isDiscount = x >= 100000
	isCashback = x >= 200000 && isCard == true

	fmt.Println("Kartu?", isCard)
	fmt.Println("Diskon?", isDiscount)
	fmt.Println("Cashback?", isCashback)

}
