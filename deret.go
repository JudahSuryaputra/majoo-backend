package main

import "fmt"

func main() {
	a := 2
	b := 4
	x := 5
	ratio := b - a
	num := a

	for i := 0; i < x; i++ {
		fmt.Print(num)
		num = num + ratio
	}
}

/* PSEUDOCODE

Deklarasi:
	var inputan deret pertama = a
	var inputan deret kedua = b
	var panjang deret = x
	var num = 0
	var pengulangan = i

Algoritma:
	rasio = b-a

	For i lebih kecil dari x
		OUTPUT num + rasio

*/
