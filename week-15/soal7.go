package main

import "fmt"

func main() {
	
	var x, j, bilangan, nx, nZero int
	
	fmt.Scan(&x)
	
	for j = 0; j < 9; j++ {
		fmt.Scan(&bilangan)
		if bilangan == x {
			nx++ 
		} else if bilangan == 0 {
			nZero++ 
		}
	}
	
	if nx > nZero {
		fmt.Printf("Modus = %d\n", x)  
	} else {
		fmt.Printf("Modus = 0\n")  
	}
}
