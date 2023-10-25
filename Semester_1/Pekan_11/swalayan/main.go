package main

import (
    "fmt"
)

func main() {
    var membership string;
	var a, b, c, d, e int;
	var isCashback, isDiscount bool;
	var cashback, discount float64;

    fmt.Scan(&membership, &a, &b, &c, &d, &e)
	fmt.Println("");

	isCashback = false;
	isDiscount = false;

	if a % 2 == 1 && b % 2 == 1 && c % 2 == 1 && d % 2 == 1 && e % 2 == 1{
		isDiscount = true;
	} else if a % 2 == 0 && b % 2 == 0 && c % 2 == 0 && d % 2 == 0 && e % 2 == 0 {
		isCashback = true;
	} else {
		isDiscount = true;
		isCashback = true;
	}

	if isCashback{
		cashback = 3.1 * float64(a + b + c)
	}

	if isDiscount {
		discount = 1.7 * float64(e + d + c)
	}

	if membership == "Yes"{
		cashback += 15
		discount += 15
	}

	if cashback > 35 {
		cashback = 35
	}

	if discount > 50 {
		discount = 50
	}

	fmt.Println("Cashback:", cashback, "% Diskon:", discount, "%")
}
