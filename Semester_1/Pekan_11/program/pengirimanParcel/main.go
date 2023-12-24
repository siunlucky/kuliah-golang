package main

import (
    "fmt"
)

func main() {
    var x, kg, gr, hargaKg, hargaGr int;

    fmt.Scan(&x)
	fmt.Println("");

	kg = x / 1000
	gr = x - kg * 1000

	hargaKg = kg * 10000
	if kg >= 10 {
		hargaGr = 0
	} else {
		if gr < 500 {
			hargaGr = gr * 15
		} else {
			hargaGr = gr * 5
		}
	}

	fmt.Println(kg, "kg +", gr, "gr")
	fmt.Println("Rp.", hargaKg, "+ Rp.", hargaGr, "= Rp.", hargaKg + hargaGr)


}
