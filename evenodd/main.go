package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range nums {
		if val%2 == 0 {
			fmt.Printf("%v is Even\n", val)
		} else {
			fmt.Printf("%v is Odd\n", val)
		}
	}
}
