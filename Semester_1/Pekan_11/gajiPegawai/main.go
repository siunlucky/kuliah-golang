package main

import (
    "fmt"
)

func main() {
    var jabatan string;
	var masaKerja, jumlahAnak, gajiPokok, tunjanganAnak int;

    fmt.Scan(&jabatan, &masaKerja, &jumlahAnak)
	fmt.Println("");

	if jumlahAnak > 3 {
		jumlahAnak = 3
	}

	if jabatan == "Staf" {
		gajiPokok = 4000
		tunjanganAnak = jumlahAnak * 100

		if masaKerja < 5 {
			tunjanganAnak *= 0
		} else if masaKerja > 10 {
			gajiPokok += 1000
		}
		
	} else if jabatan == "Manager" {
		tunjanganAnak = jumlahAnak * 300
		gajiPokok = 8500

		if masaKerja > 10 {
			gajiPokok += 1500
		}

	} else if jabatan == "Direktur" {
		gajiPokok = 20000
		tunjanganAnak = jumlahAnak * 500
	}

	fmt.Println(gajiPokok, "+", tunjanganAnak, "=", gajiPokok + tunjanganAnak)


}
