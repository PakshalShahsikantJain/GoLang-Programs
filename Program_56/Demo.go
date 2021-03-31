package main

import "fmt"

func EvenFactorial(No int) int {
	Value := 1

	for i:=No ;i > 0;i--{
		if i % 2 == 0 {
			Value = Value * i
		}
	}

	return Value
}
func main() {
	No := 0
	ret := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = EvenFactorial(No)
	fmt.Println("Even Factorial of Given No is :",ret)
}