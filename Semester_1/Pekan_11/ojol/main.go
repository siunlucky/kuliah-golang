package main

import (
    "fmt"
)

func main() {
    var h, m, harga int;
	var km float64;

	fmt.Scan(&h, &m, &km)
	fmt.Println("")

	if h >= 5 && (h < 22 || h == 22 && m == 0) {
		if (h >= 5 && (h < 9 || h == 9 && m == 0)) || (h >= 16 && (h < 19 || h == 19 && m == 0)) {
			if km > 0 && km <= 10 {
				harga = int(km * 5000)
			} else if km > 10 && km <= 20 {
				harga = int(km * 4500)
			} else {
				fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
			}
		} else {
			if km > 0 && km <= 20 {
				harga = int(km * 4000)
			} else {
				fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
			}
		}
		fmt.Println(harga)
	} else {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
	}

}
