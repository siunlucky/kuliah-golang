package main

import (
	"fmt"
    "math"
    "strings"
    "os"
    "bufio"
);

func main(){
    var soal int

    fmt.Println("")
    fmt.Println("Nama : Faiz Elfahad Kurniawan")
    fmt.Println("NIM : 103012300408")
    fmt.Println("Kelas : IF-47-11")

    for {
        fmt.Println("")
        fmt.Println("=========== PILIH TUGAS SOAL ===========")
        fmt.Println("")
        fmt.Println("1.  Menghitung luas lingkaran")
        fmt.Println("2.  Menghitung diameter lingkaran")
        fmt.Println("3.  Mengkonversi suhu dari Celsius menjadi Reamur")
        fmt.Println("4.  Menghitung keliling persegi panjang")
        fmt.Println("5.  Mengkonversi suhu dari Celsius menjadi Fahrenheit")
        fmt.Println("6.  Mengkonversi suhu dari Celsius menjadi Kelvin")
        fmt.Println("7.  Mengkonversi suhu dari Celsius menjadi Kelvin")
        fmt.Println("8.  Menghitung luas permukaan bola")
        fmt.Println("9.  Menentukan bilangan ganjil atau bukan")
        fmt.Println("10. Menentukan bilangan bulat bilangan negatif atau bukan")
        fmt.Println("11. Menentukan bilangan bulat bilangan positif atau bukan")
        fmt.Println("12. Menentukan huruf hidup (vokal) atau bukan")
        fmt.Println("13. Menentukan huruf mati (konsonan) atau bukan")
        fmt.Println("14. Menentukan huruf mati (konsonan) atau bukan")
        fmt.Println("15. Menentukan bilangan nol atau bukan")
        fmt.Println("16. Mencetak String Nama dengan format")
        fmt.Println("17. Mencetak karakter")
        fmt.Println("0.  EXIT")
    
        fmt.Println("")
        fmt.Print("Masukkan Nomor Soal: ")
        fmt.Scan(&soal)
        fmt.Println("")

        switch soal {
        case 1:
            luasLingkaran();
        case 2:
            diameterLingkaran();
        case 3:
            celciusToReamur();
        case 4:
            kelilingPersegiPanjang();
        case 5:
            celciusToFahrenheit();
        case 6, 7:
            celciusToKelvin();
        case 8: 
            luasPermukaanBola();
        case 9:
            menentukanBilanganGanjil();
        case 10:
            menentukanBilanganNegatif();
        case 11:
            menentukanBilanganPositif();
        case 12:
            menentukanHurufVokal();
        case 13, 14:
            menentukanHurufKonsonan();
        case 15:
            menentukanBilanganNol();
        case 16:
            mencetakStringNama();
        case 17:
            mencetakStringKarakter();
        case 0:
            os.Exit(1)
        default:
            continue
        }
        
        

        fmt.Println("");
        fmt.Println("Lanjutkan? (yes/no)")

        var lanjut string

        fmt.Scan(&lanjut)

        if lanjut == "yes" {
            continue
        } else {
            fmt.Println("Author : Faiz Elfahad Kurniawan")
            break
        }
    }
}

// Soal 1

func luasLingkaran(){
    var pi float64
    var r float64

	pi = 3.14

    fmt.Println("=========== SOAL NO 1 ===========")
    fmt.Println("Menghitung luas lingkaran")
    fmt.Println("")
    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scan(&r)

    luas := pi * r * r

    fmt.Printf("Luas lingkaran adalah %f\n", luas)

}

// Soal 2

func diameterLingkaran(){
    var pi float64
    var luas float64
    
	pi = 3.14

    fmt.Println("=========== SOAL NO 2 ===========")
    fmt.Println("Menghitung diameter lingkaran")
    fmt.Println("")
    fmt.Print("Masukkan luas lingkaran: ")
    fmt.Scan(&luas)

    r := math.Sqrt(luas / pi)
    d := 2 * r
    
    fmt.Printf("Diameter Lingkaran adalah %f\n", d)
}

// Soal 3

func celciusToReamur(){
    var celcius float64

    fmt.Println("=========== SOAL NO 3 ===========")
    fmt.Println("Mengkonversi suhu dari derajat Celsius menjadi derajat Reamur")
    fmt.Println("")
    fmt.Print("Masukkan nilai Celcius: ")
    fmt.Scan(&celcius)

    reamur := 4.0 / 5.0 * celcius

    fmt.Printf("Hasil dari konversi suhu dari celcius ke reamur adalah %f\n", reamur)
}

// Soal 4

func kelilingPersegiPanjang(){
    var panjang, lebar int

    fmt.Println("=========== SOAL NO 4 ===========")
    fmt.Println("Menghitung Keliling Persegi Panjang")
    fmt.Println("")
    fmt.Print("Masukkan panjang: ")
    fmt.Scan(&panjang)
    fmt.Print("Masukkan lebar: ")
    fmt.Scan(&lebar)

    keliling := (panjang + lebar) * 2

    fmt.Println("Keliling persegi adalah", keliling)
}

// Soal 5

func celciusToFahrenheit(){
    var celcius float64

    fmt.Println("=========== SOAL NO 5 ===========")
    fmt.Println("Mengkonversi suhu dari derajat Celsius menjadi derajat Fahrenheit")
    fmt.Println("")
    fmt.Print("Masukkan nilai Celcius: ")
    fmt.Scan(&celcius)

    fahrenheit := (9.0 / 5.0) * celcius + 32 

    fmt.Printf("Hasil dari konversi suhu dari celcius ke Fahrenheit adalah %f\n", fahrenheit)
}

// Soal 6 & 7

func celciusToKelvin(){
    var celcius float64

    fmt.Println("=========== SOAL NO 6 & 7 ===========")
    fmt.Println("Mengkonversi suhu dari derajat Celsius menjadi derajat Kelvin")
    fmt.Println("")
    fmt.Print("Masukkan nilai Celcius: ")
    fmt.Scan(&celcius)

    kelvin := celcius + 273

    fmt.Printf("Hasil dari konversi suhu dari Celcius ke Kelvin adalah %f\n", kelvin)
}

// Soal 8

func luasPermukaanBola(){
    var r float64
    var pi float64

    pi = 3.14

    fmt.Println("=========== SOAL NO 8 ===========")
    fmt.Println("Menghitung luas permukaan bola")
    fmt.Println("")
    fmt.Print("Masukkan nilai jari-jari bola: ")
    fmt.Scan(&r)

    luasPermukaanBola := 4.0 * pi * (r * r)

    fmt.Printf("Luas Permukaan Bola %f\n", luasPermukaanBola)
}

// Soal 9

func menentukanBilanganGanjil(){
    var x int
    var isOdd bool

    fmt.Println("=========== SOAL NO 9 ===========")
    fmt.Println("Menentukan apakah bilangan ganjil atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&x)

    isOdd = (x % 2) == 1

    fmt.Printf("%t", isOdd)
}

// Soal 10

func menentukanBilanganNegatif(){
    var x int
    var isNegative bool

    fmt.Println("=========== SOAL NO 10 ===========")
    fmt.Println("Menentukan apakah bilangan negatif atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&x)

    isNegative = x < 0

    fmt.Printf("%t", isNegative)
}

// Soal 11

func menentukanBilanganPositif(){
    var x int
    var isPositive bool

    fmt.Println("=========== SOAL NO 11 ===========")
    fmt.Println("Menentukan apakah bilangan positif atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&x)

    isPositive = x > 0

    fmt.Printf("%t", isPositive)
}

// Soal 12

func menentukanHurufVokal(){
    var x string
    var isVocal bool

    fmt.Println("=========== SOAL NO 12 ===========")
    fmt.Println("Menentukan apakah huruf vokal atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan huruf: ")
    fmt.Scan(&x)

    x  = strings.ToLower(x)

    isVocal = x == "a" || x == "e" || x == "i" || x == "o" || x == "u"

    fmt.Printf("%t", isVocal)
}

// Soal 13 & 14

func menentukanHurufKonsonan(){
    var x string
    var isConsonant bool

    fmt.Println("=========== SOAL NO 13 & 14 ===========")
    fmt.Println("Menentukan apakah huruf mati (konsonan) atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan huruf: ")
    fmt.Scan(&x)

    x  = strings.ToLower(x)

    isConsonant = x != "a" && x != "e" && x != "i" && x != "o" && x != "u"

    fmt.Printf("%t", isConsonant)
}

// Soal 15

func menentukanBilanganNol(){
    var x int
    var isZero bool

    fmt.Println("=========== SOAL NO 15 ===========")
    fmt.Println("Menentukan apakah bilangan nol atau bukan")
    fmt.Println("")
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&x)

    isZero = x == 0

    fmt.Printf("%t", isZero)
}

// Soal 16

func mencetakStringNama() {
    fmt.Println("=========== SOAL NO 16 ===========")
    fmt.Println("Mencetak string nama")
    fmt.Println("")
    fmt.Print("Masukkan Nama: ")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        name := scanner.Text()
        if name != "" {
            fmt.Println("")
            fmt.Print("Hey, ", name)
            fmt.Println("")
            break
        }
        fmt.Print(name)
    }

    if err := scanner.Err(); err != nil {
        os.Exit(1)
    }
    
}

// Soal 17

func mencetakStringKarakter() {
    var karakter rune

    fmt.Println("=========== SOAL NO 17 ===========")
    fmt.Println("Mencetak Karakter")
    fmt.Println("")
    fmt.Print("Masukkan Karakter: ")
    fmt.Scanf("%c\n", &karakter)
    
    fmt.Printf("Karakter yang Anda ketik adalah: %c\n", karakter)
}

