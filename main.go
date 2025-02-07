package main

import "fmt"

func getHello(text string) string {
	return text
}

func pertambahan(angka1, angka2 int) int {
	return angka1 + angka2
}

func main() {
	fmt.Println(getHello("hello world"))
	fmt.Println(pertambahan(4, 5))
}
