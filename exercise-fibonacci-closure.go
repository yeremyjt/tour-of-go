package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	j := 1
	counter := 0
	return func() int {
		if (counter == 0) {
			counter++
			return 0
		}
		if (counter == 1) {
			counter++
			return 1
		}
		sum := i + j
		i = j
		j = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
