package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangele() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
