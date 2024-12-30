package main

import "fmt"

func main() {

	var input string

	fmt.Scan(&input)

	if input == "A" || input == "I" || input == "U" || input == "E" || input == "0" {
		fmt.Println("Bukan konsonan")
		return
	}
	fmt.Println("Huruf konsonan")
}