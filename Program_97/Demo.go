package main 

import "fmt"

func Kilometer(No int) int {
	Value := 0

	if(No < 100){
		Value = No * 25
	} else if No >= 100{
		Value = No * 15
	}

	return Value
}
func main(){
	No := 0
	ret := 0

	fmt.Println("Jay Ganesh")
	fmt.Println("Enter Your Run of KM of Car")
	fmt.Scan(&No)

	ret = Kilometer(No)
	fmt.Println("Your Charge for Driving Car is :",ret)
}