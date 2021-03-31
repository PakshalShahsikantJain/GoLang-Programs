package main

import "fmt"

func Mult(No, No2, No3 int) int {
	if No < 0 || No2 < 0 || No3 < 0 {
		No = -No
		No2 = -No2
		No3 = -No3
	}

	if No == 0 {
		No = 1
	} else if No2 == 0 {
		No2 = 1
	} else if No3 == 0 {
		No3 = 1
	}

	Mult := 0

	Mult = No * No2 * No3

	return Mult
}

func main() {
	fmt.Println("Jay Ganesh.....")

	No := 0
	No2 := 0
	No3 := 0

	ret := 0

	fmt.Println("Enter Three Number's to Multiply")
	fmt.Scanf("%d %d %d", &No, &No2, &No3)

	ret = Mult(No, No2, No3)
	fmt.Println("Multiplication is : ", ret)
}
