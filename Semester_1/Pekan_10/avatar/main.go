package main

import "fmt"

func main() {
    var jumlahOrang, besar, kecil int
    fmt.Print("Masukkan jumlah orang: ")
    fmt.Scan(&jumlahOrang)

    for jumlahOrang >= 5 && besar < 3 {
        besar++
        jumlahOrang -= 5
    }

    if jumlahOrang >= 1 && besar < 3 {
        besar++
        jumlahOrang -= jumlahOrang
    }

    for jumlahOrang >= 2 && kecil < 5 {
        kecil++
        jumlahOrang -= 2
    }

    if jumlahOrang >= 1 && kecil < 5 {
        kecil++
        jumlahOrang -= jumlahOrang
    }

    if besar > 0 {
        fmt.Print("Dewasa ", besar)
    }

    if kecil > 0 {
        fmt.Print(", kecil ", kecil)
    }

    if jumlahOrang > 0 {
        fmt.Print(", dan ", jumlahOrang, " tak berangkat ")
    }

}
