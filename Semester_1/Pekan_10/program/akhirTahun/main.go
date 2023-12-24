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
	isCashback = x >= 200000 && isCard
	isCard = isCard && isDiscount

	if isDiscount {
		x -= (x * 10) / 100 // x dikurangi discount 10%
	}

	if isCashback {
		x -= 75000 // x dikurangi cashback 75000
	}

	fmt.Println("Kartu?", isCard)
	fmt.Println("Diskon?", isDiscount)
	fmt.Println("Cashback?", isCashback)
	fmt.Println("Rp.", x)

}
