package main

import "fmt"

func DisplayFactors(No int) {
	i := 0

	for i = No / 2; i > 0; i-- {
		if No%i == 0 {
			fmt.Print(i, "\t")
		}
	}
}

func main() {
	fmt.Println("Jay Ganesh...")
	No := 0
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &No)

	DisplayFactors(No)
}
