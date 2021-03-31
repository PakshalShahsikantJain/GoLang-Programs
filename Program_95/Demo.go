package main

import "fmt"

func Tax(No int) {
	Value := 0
	if No < 500000 {
		return
	} else if No >= 500000 && No <= 1000000 {
		Value = ((No / 100 )* 10)
		fmt.Println(Value)
	} else if No >= 1000000 && No <=2000000 {
		Value = ((No / 100) * 20)
		fmt.Println(Value)
	} else {
		Value = ((No / 100) * 30)
		fmt.Println(Value)
	}	
}
func main() {
	No := 0
	No2 := 0
	//ret := 0

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Your Annual Income")
	fmt.Scan(&No)

	No2 = No / 12

	Tax(No2)
	//fmt.Printf("Your Taxable Amount is %d:",ret)
}