package main

import "fmt"

func MultFactors(No int) int {
	i := 0
	Mult := 1

	for i = 1; i <= No/2; i++ {
		if No%i == 0 {
			Mult = Mult * i
		}
	}

	return Mult
}

func main() {
	fmt.Println("Jay Ganesh......")

	No := 0
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &No)

	ret := MultFactors(No)
	fmt.Printf("Multiplication of Factors is : %d \n", ret)
}
