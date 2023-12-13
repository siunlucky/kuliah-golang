package main

import "fmt";

func main() {
	var e0, c, e1, hasil int
	
	fmt.Scan(&e0, &c, &e1);

	hasil = (e0 - e1) / c

	fmt.Println(hasil)
}