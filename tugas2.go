package main

import "fmt"

func main() {
	const angka = 101

	if angka%2 == 0 {
		fmt.Println("Nilai Genap")
	} else {
		fmt.Println("Nilai Ganjil")
	}
}
