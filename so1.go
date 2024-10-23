package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Masukkan bilangan x :")
	fmt.Scanln(&x)

	fx := 2/(x+5) + 5 //float

	fmt.Println("Nilai dari x adalah ", (fx))
}
