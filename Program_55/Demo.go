package main

import "fmt"

func DollarToINR(No int) int {

	if No < 0{
		No = -No
	}
	Value := 0

	Value = No * 70 // As 1 Dollar is Equal to 70 rupees

	return Value
}
func main(){
	No := 0
	ret := 0
	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Amount In US Dollar")
	fmt.Scan(&No)

	ret = DollarToINR(No)

	fmt.Println("Conversion of amount Dollar To  INR is :",ret)
}