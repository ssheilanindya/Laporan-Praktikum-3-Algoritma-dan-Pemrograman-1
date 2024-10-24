package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const pi = 3.1415926535

	fmt.Print("Masukkan jari-jari bola :")
	fmt.Scan(&r)

	volumeBola := (4.0 / 3.0) * pi * math.Pow(r, 3)
	luasBola := 4 * pi * math.Pow(r, 2)

	fmt.Printf("Volume bola adalah %.2f\n", volumeBola)
	fmt.Printf("Luas permukaan bola adalah %.2f\n", luasBola)

}
