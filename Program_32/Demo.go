package main

import "fmt"

func Mult(No int) int {
	Value := 0
	Mult := 1

	if No < 0 {
		No = -No
	}

	for No != 0 {
		Value = No % 10
		if Value == 0 {
			Value = 1
		}
		Mult = Mult * Value
		No = No / 10
	}

	return Mult
}
func main() {
	fmt.Println("Jay Ganesh......")
	No := 0
	ret := 0

	fmt.Println("Enter Number..")
	fmt.Scan(&No)

	ret = Mult(No)
	fmt.Println("Multiplication of Digits is : ", ret)
}
