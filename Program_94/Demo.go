package main 

import "fmt"

func Discount(No int){
	Value := 0
	Value2 := 0
	TValue1 := 0
	TValue2 := 0

	if No >= 500 && No <= 1500 {
		fmt.Println("You Get Discount of 10 % :")
		Value = ((No/100) * 10)
		TValue1 = No - Value
		fmt.Println("Your Final Bill Amount after Discount is :",TValue1)
	} else if No > 1500 {
		fmt.Println("You Get Discount of 15 % :")
		Value2 = ((No/100) * 15)
		TValue2 = No - Value2
		fmt.Println("Your Final Bill Amount after Discount is :",TValue2) 
	} else {
		fmt.Println("You Don't Get Any Discount Your Final Bill Amount is :",No)
	}
}
func main() {
	No := 0

	fmt.Println("Jay Ganesh...")
	fmt.Println("Welcome to Hotel V,S Karande")
	fmt.Println("Enter Your Total Bill Amount")
	fmt.Scan(&No)

	Discount(No)
}