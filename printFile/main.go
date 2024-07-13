package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Args[1])

	// f, err := os.Open(os.Args[1])

	// if err !=  nil {
	// 	fmt.Println("Error : ", err)
	// 	os.Exit(1)
	// }

	// io.Copy(os.Stdout, f)

	f , err := os.ReadFile("myFile.txt");

	if err !=  nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	fmt.Println(string(f)) //... 2
}