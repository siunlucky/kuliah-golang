package main

import (
	"fmt"
);

func main()  {
	var attempts int;
	var username, password string;

	attempts = -1;
	for username != "admin" && password != "admin" {
		fmt.Scan(&username, &password)
		attempts++
	}

	fmt.Println(attempts);
	fmt.Println("Login berhasil");

}