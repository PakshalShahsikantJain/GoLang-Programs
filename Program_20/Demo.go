package main

import "fmt"

func Difference(No int) int {
	i := 0
	j := 0
	Sum1 := 0
	Sum2 := 0

	for i = 1; i <= No/2; i++ {
		if No%i == 0 {
			Sum1 = Sum1 + i
		}
	}

	for j = 1; j < No; j++ {
		if No%j > 0 {
			Sum2 = Sum2 + j
		}
	}

	return Sum1 - Sum2
}

func main() {
	fmt.Println("Jay Ganesh")
	No := 0
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &No)

	ret := Difference(No)
	fmt.Println("Difference is : ", ret)
}
