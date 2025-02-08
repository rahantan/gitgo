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

type Dosen struct {
	Name, Matkul string
}

func GetDataDosen() map[string]any {

	return map[string]any{
		"1": Dosen{Name: "joko1", Matkul: "pemograman1"},
		"2": Dosen{Name: "joko2", Matkul: "pemograman2"},
		"3": Dosen{Name: "joko3", Matkul: "pemograman3"},
	}
}

type Staff struct {
	Name, Age string
}

func GetDataStaff() map[string]any {
	return map[string]any{
		"1": Staff{Name: "staf1", Age: "20"},
		"2": Staff{Name: "staf2", Age: "20"},
		"3": Staff{Name: "staf3", Age: "20"},
	}
}
func main() {
	dosen := GetDataDosen()
	for _, value := range dosen {
		fmt.Println(value)
		if data, ok := value.(Dosen); ok {
			fmt.Printf("nama: %s matakuliah %s\n", data.Name, data.Matkul)
		}

	}
	staff := GetDataStaff()
	for _, value := range staff {
		if data, ok := value.(Staff); ok {
			fmt.Printf("Name: %s Age %s\n", data.Name, data.Age)
		}
	}

	fmt.Println(getName("sukma"))
}
