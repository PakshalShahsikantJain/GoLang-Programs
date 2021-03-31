package main

import "fmt"

func DisplayNonFactors(No int) {
	i := 0

	for i = 2; i < No; i++ {
		if No%i > 0 {
			fmt.Print(i, "\t")
		}
	}
}
func main() {
	fmt.Println("Jay Ganesh")
	No := 0
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &No)

	DisplayNonFactors(No)
}
