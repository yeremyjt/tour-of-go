package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tempz := 1.0
	diff := math.MaxFloat64
	for diff > 0.0000000000000001  {
		tempz -= (tempz*tempz - x) / (2*tempz)
		diff = math.Abs(tempz - z)
		z = tempz
		fmt.Printf("z = %g, diff = %g\n", tempz, diff)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}