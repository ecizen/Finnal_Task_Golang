package main

import "fmt"

func main() {
	var T, totalAir, V int


	fmt.Scan(&T)

	totalAir = 0
	


	for {
		fmt.Scan(&V)
		
		totalAir += V

		if totalAir >= T {
			fmt.Println("true")
			break
		} else {
			fmt.Println("false")
		}
	}
}
