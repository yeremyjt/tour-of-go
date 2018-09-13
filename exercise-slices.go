package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := range a {
		a[i] = make([]uint8, dx)
	}
	
	// Do something.
	for i := range a {
		for j := range a[i] {
			switch {
			case j % 15 == 0:
				a[i][j] = uint8((i+j)/2)
			case j % 3 == 0:
				a[i][j] = uint8(i*j)
			case j % 5 == 0:
				a[i][j] = uint8(i^j)
			default:
				a[i][j] = 100
			}
		}
	}	
	return a
}

func main() {
	pic.Show(Pic)
}

