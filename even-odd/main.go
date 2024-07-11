package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	numbers := getNumbers(5)

	for _, val := range numbers {
		if val%2 == 0 {
			fmt.Println(val, " Number is Even")
		} else {
			fmt.Println(val, " Number is Odd")
		}
	}
}