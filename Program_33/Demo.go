package main

import "fmt"

func Difference(No int) int {
	Value := 0
	SumEven := 0
	SumOdd := 0

	for No != 0 {
		Value = No % 10
		if Value%2 == 0 {
			SumEven = SumEven + Value
		} else if Value%2 == 1 {
			SumOdd = SumOdd + Value
		}

		No = No / 10
	}

	return SumEven - SumOdd
}

func main() {
	fmt.Println("Jay Ganesh....")

	No := 0
	ret := 0
	fmt.Println("Enter Number..")
	fmt.Scan(&No)

	ret = Difference(No)
	fmt.Println("Difference Between Even and Odd Number is : ", ret)
}
