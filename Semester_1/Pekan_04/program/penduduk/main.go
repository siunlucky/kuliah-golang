package main

import "fmt";

func main() {
	var jumlahPendudukAwal, jumlahKelahiran, jumlahImigrasi, jumlahKematian, jumlahEmigrasi, hasil int

	fmt.Scan(&jumlahPendudukAwal, &jumlahKelahiran, &jumlahImigrasi, &jumlahKematian, &jumlahEmigrasi)
	hasil = jumlahPendudukAwal + jumlahKelahiran + (jumlahImigrasi - (jumlahKematian + jumlahEmigrasi))

	fmt.Println(hasil)
}