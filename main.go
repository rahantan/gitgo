package main

import (
	"fmt"
	"strings"
)

func hitungFrekuensiPerkata() map[string]int {
	// algoritma
	input := "hello teman semuanya saya, cuman memberi tahu teman teman"

	// Konversi teks ke huruf kecil → agar tidak membedakan huruf besar/kecil.
	// Pisahkan kata → menggunakan strings.Fields().
	filter := strings.Fields(strings.ToLower(input)) // Pisahkan kata

	// Gunakan map → untuk menyimpan jumlah kemunculan setiap kata.
	data := make(map[string]int)

	for _, f := range filter {
		data[f]++
	}
	// Cetak hasilnya → tampilkan kata dan jumlah kemunculannya.
	return data
}

func GanjilGenap() ([]int, []int) {
	angka := []int{
		5, 6, 2, 1, 6, 8, 3, 2, 1, 2, 5,
	}
	var ganjil, genap []int
	for _, a := range angka {
		if a%2 == 0 {
			ganjil = append(ganjil, a)
		} else {
			genap = append(genap, a)
		}
	}
	return ganjil, genap
}

func Rekursif(data *[]int, angka int) {
	if angka < 0 {
		return
	}
	Rekursif(data, angka-1)
	*data = append(*data, angka)
}
func main() {
	var data []int
	Rekursif(&data, 4)
	fmt.Println(data)
}
