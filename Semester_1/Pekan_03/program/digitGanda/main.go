package main

import "fmt";

func main() {
	var x, digitPertama, digitKedua, hasil int

	fmt.Scan(&x)
	
	digitPertama = x / 10
	digitKedua = x % 10
	
	hasil = digitPertama * 1000 + digitPertama * 100 + digitKedua * 10 + digitKedua

	fmt.Println(hasil)
}