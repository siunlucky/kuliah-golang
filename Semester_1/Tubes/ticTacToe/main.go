package main

import "fmt"

func main() {
	var A1, A2, A3, B1, B2, B3, C1, C2, C3, namePlayer1, namePlayer2, rolePlayer1, rolePlayer2, kotak, last, retry string
	var isRetry bool
	var countMatch, winPlayer1, winPlayer2 int

	rolePlayer1 = " "
	isRetry = true
	countMatch, winPlayer1, winPlayer2 = 0, 0, 0

	fmt.Print("Ketik Nama Player 1 : ")
	fmt.Scan(&namePlayer1)

	fmt.Print("Ketik Nama Player 2 : ")
	fmt.Scan(&namePlayer2)
	fmt.Println("")

	fmt.Println("Hai", namePlayer1, "dan", namePlayer2, ", Selamat datang di permainan Tic Tac Toe!!")
	fmt.Println("")

	fmt.Println(namePlayer1, "silahkan memilih ingin main sebagai X atau O")
	fmt.Println("")
	for rolePlayer1 == " " {
		fmt.Print("X / O : ")
		fmt.Scan(&rolePlayer1)

		if rolePlayer1 == "X" || rolePlayer1 == "x" {
			rolePlayer2 = "O"
		} else if rolePlayer1 == "O" || rolePlayer1 == "o" {
			rolePlayer2 = "X"
		} else {
			rolePlayer1 = " "
			fmt.Println("Input yang dimasukkan salah! Silahkan coba lagi")
		}
	}

	for isRetry {
		if countMatch > 0 {
			if rolePlayer1 == "X" {
				rolePlayer1 = "O"
				rolePlayer2 = "X"
			} else {
				rolePlayer1 = "X"
				rolePlayer2 = "O"
			}
		}

		if rolePlayer1 == "X" {
			last = namePlayer1
		} else {
			last = namePlayer2
		}

		A1, A2, A3, B1, B2, B3, C1, C2, C3 = " ", " ", " ", " ", " ", " ", " ", " ", " "

		fmt.Println("")
		fmt.Println(namePlayer1, "adalah", rolePlayer1)
		fmt.Println(namePlayer2, "adalah", rolePlayer2)
		fmt.Println("")

		fmt.Println("   ", "A", " ", "B", " ", "C")
		fmt.Println("1", "|", A1, "|", B1, "|", C1, "|")
		fmt.Println("2", "|", A2, "|", B2, "|", C2, "|")
		fmt.Println("3", "|", A3, "|", B3, "|", C3, "|")

		fmt.Println("")
		fmt.Println("PERMAINAN DIMULAI!!")
		fmt.Println("")

		for (A1 == " " || A2 == " " || A3 == " " || B1 == " " || B2 == " " || B3 == " " || C1 == " " || C2 == " " || C3 == " ") &&
			(!((A1 == "X" && A2 == "X" && A3 == "X") || (A1 == "O" && A2 == "O" && A3 == "O")) &&
				!((B1 == "X" && B2 == "X" && B3 == "X") || (B1 == "O" && B2 == "O" && B3 == "O"))) &&
			!((C1 == "X" && C2 == "X" && C3 == "X") || (C1 == "O" && C2 == "O" && C3 == "O")) &&
			!((A1 == "X" && B1 == "X" && C1 == "X") || (A1 == "O" && B1 == "O" && C1 == "O")) &&
			!((A2 == "X" && B2 == "X" && C2 == "X") || (A2 == "O" && B2 == "O" && C2 == "O")) &&
			!((A3 == "X" && B3 == "X" && C3 == "X") || (A3 == "O" && B3 == "O" && C3 == "O")) &&
			!((A1 == "X" && B2 == "X" && C3 == "X") || (A1 == "O" && B2 == "O" && C3 == "O")) &&
			!((C1 == "X" && B2 == "X" && A3 == "X") || (C1 == "O" && B2 == "O" && A3 == "O")) {
			if last == namePlayer1 {
				fmt.Print("Selanjutnya, ", namePlayer1, " silahkan memilih kotak yang ingin diisi : ")
				fmt.Scan(&kotak)
				if kotak == "A1" || kotak == "a1" {
					if A1 == " " {
						A1 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "A2" || kotak == "a2" {
					if A2 == " " {
						A2 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "A3" || kotak == "a3" {
					if A3 == " " {
						A3 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B1" || kotak == "b1" {
					if B1 == " " {
						B1 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B2" || kotak == "b2" {
					if B2 == " " {
						B2 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B3" || kotak == "b3" {
					if B3 == " " {
						B3 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C1" || kotak == "c1" {
					if C1 == " " {
						C1 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C2" || kotak == "c2" {
					if C2 == " " {
						C2 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C3" || kotak == "c3" {
					if C3 == " " {
						C3 = rolePlayer1
						last = namePlayer2
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else {
					fmt.Println("")
					fmt.Println("Kotak tidak ditemukan. Silahkan memilih kotak yang lain")
				}
			} else {
				fmt.Print("Selanjutnya, ", namePlayer2, " silahkan memilih kotak yang ingin diisi : ")
				fmt.Scan(&kotak)
				if kotak == "A1" || kotak == "a1" {
					if A1 == " " {
						A1 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "A2" || kotak == "a2" {
					if A2 == " " {
						A2 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "A3" || kotak == "a3" {
					if A3 == " " {
						A3 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B1" || kotak == "b1" {
					if B1 == " " {
						B1 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B2" || kotak == "b2" {
					if B2 == " " {
						B2 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "B3" || kotak == "b3" {
					if B3 == " " {
						B3 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C1" || kotak == "c1" {
					if C1 == " " {
						C1 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C2" || kotak == "c2" {
					if C2 == " " {
						C2 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else if kotak == "C3" || kotak == "c3" {
					if C3 == " " {
						C3 = rolePlayer2
						last = namePlayer1
					} else {
						fmt.Println("")
						fmt.Println("Kotak telah terisi! Silahkan memilih kotak yang lain")
					}
				} else {
					fmt.Println("")
					fmt.Println("Kotak tidak ditemukan. Silahkan memilih kotak yang lain")
				}
			}

			fmt.Println("")
			fmt.Println("   ", "A", " ", "B", " ", "C")
			fmt.Println("1", "|", A1, "|", B1, "|", C1, "|")
			fmt.Println("2", "|", A2, "|", B2, "|", C2, "|")
			fmt.Println("3", "|", A3, "|", B3, "|", C3, "|")
			fmt.Println("")
		}

		if (A1 == "X" && A2 == "X" && A3 == "X") || (B1 == "X" && B2 == "X" && B3 == "X") || (C1 == "X" && C2 == "X" && C3 == "X") || (A1 == "X" && B2 == "X" && C3 == "X") || (C1 == "X" && B2 == "X" && A3 == "X") || (A1 == "X" && B1 == "X" && C1 == "X") || (A2 == "X" && B2 == "X" && C2 == "X") || (A3 == "X" && B3 == "X" && C3 == "X") {
			if rolePlayer1 == "X" {
				fmt.Println(namePlayer1, "Memenangkan Pertandingan!!")
				winPlayer1++
			} else {
				fmt.Println(namePlayer2, "Memenangkan Pertandingan!!")
				winPlayer2++
			}
		} else if (A1 == "O" && A2 == "O" && A3 == "O") || (B1 == "O" && B2 == "O" && B3 == "O") || (C1 == "O" && C2 == "O" && C3 == "O") || (A1 == "O" && B2 == "O" && C3 == "O") || (C1 == "O" && B2 == "O" && A3 == "O") || (A1 == "O" && B1 == "O" && C1 == "O") || (A2 == "O" && B2 == "O" && C2 == "O") || (A3 == "O" && B3 == "O" && C3 == "O") {
			if rolePlayer1 == "O" {
				fmt.Println(namePlayer1, "Memenangkan Pertandingan!!")
				winPlayer1++
			} else {
				fmt.Println(namePlayer2, "Memenangkan Pertandingan!!")
				winPlayer2++
			}
		} else {
			fmt.Println("Pertandingan Seri!!")
		}

		fmt.Println("")
		fmt.Println("---------------- SCORE ----------------")
		fmt.Println(namePlayer1, ":", winPlayer1)
		fmt.Println(namePlayer2, ":", winPlayer2)
		fmt.Println("")

		for retry != "N" && retry != "Y" {
			fmt.Print("Bermain Lagi ? Y/N : ")
			fmt.Scan(&retry)
		}

		if retry == "Y" {
			isRetry = true
			fmt.Println("")
			fmt.Println("Posisi Ditukar!")
		} else if retry == "N" {
			isRetry = false
		}
		retry = ""
		countMatch++
	}

	fmt.Println(" ")
	fmt.Println("Created by :")
	fmt.Println("Faiz Elfahad Kurniawan")
	fmt.Println("103012300408")
}
