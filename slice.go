package main

import (
	"fmt"
)

func main() {
	// Membuat sebuah slice dengan jumlah data awal 5
	slice := make([]int, 5)

	// Menambahkan data ke dalam slice sebanyak 3
	for i := 1; i <= 3; i++ {
		// Menambahkan data ke dalam slice menggunakan append
		slice = append(slice, i*10)
	}

	// Menampilkan slice
	fmt.Println("Slice setelah menambahkan data:")
	fmt.Println(slice)
}
