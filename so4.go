package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Masukkan suhu dalam derajat fahrenheit :")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	reamur := celsius * 4 / 5
	kelvin := (fahrenheit + 459.67) * 5 / 9

	fmt.Printf("Suhu dalam Celsius adalah %.2f\n", celsius)
	fmt.Printf("Suhu dalam Reamur adalah %.2f\n", reamur)
	fmt.Printf("Suhu dalam Kelvin adalah %.2f\n", kelvin)
}
