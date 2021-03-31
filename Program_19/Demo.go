package main

import "fmt"

func SumOfNonFactors(No int) int {
	i := 0
	Sum := 0

	for i = 2; i < No; i++ {
		if No%i > 0 {
			Sum = Sum + i
		}
	}
	return Sum
}

func main() {
	fmt.Println("Jay Ganesh....")

	No := 0
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &No)

	ret := SumOfNonFactors(No)

	fmt.Println("Addition of Non Factors is : ", ret)
}
