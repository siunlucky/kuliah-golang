package main

import "fmt"

func main() {
    var gigi int
    var isClutchPulled, isPedalPressed bool

    fmt.Scan(&gigi, &isClutchPulled, &isPedalPressed)
    fmt.Println()

    if !isClutchPulled && !isPedalPressed {
        fmt.Println("Motor Mati")
    } else if isClutchPulled {
        fmt.Println("Mesin menyala dan motor tidak berjalan")
    } else if !isClutchPulled && gigi > 0 && isPedalPressed {
        fmt.Println("Motor berjalan")
    } else {
        fmt.Println("Mesin menyala dan motor tidak berjalan")
    }
}
