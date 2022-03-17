package main

import "fmt"

func main() {
	var s string
	var a, b, hasil_penjumlahan int

	fmt.Scanln(&s, &a, &b)
	hasil_penjumlahan = a + b

	fmt.Println("Kata -", s)
	fmt.Println("Jumlah =", hasil_penjumlahan)
}
