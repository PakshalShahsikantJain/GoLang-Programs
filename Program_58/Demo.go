package main

import "fmt"

func Difference(No int) int {
	Value := 1
	Value2 := 1

	if No < 0 {
		No = -No
	}

	for i := No ;i > 0;i-- {
		if((i % 2) == 0) { 
			Value = Value * i
		} else {
			Value2 = Value2 * i
		}
	}

	return Value - Value2
}

func main() {
	No := 0
	ret := 0
	fmt.Println("Jay Ganesh........")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Difference(No)

	fmt.Println("Difference Between Even and ODD Factorial is :",ret)
}








































