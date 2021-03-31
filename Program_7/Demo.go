package main

import "fmt"

func Check(No int) bool {

	if No/5 == 0 {
		return true
	} else {
		return false
	}
}
func main() {

	var No int
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	bRet := Check(No)

	if bRet == true {
		fmt.Println("Number is Divisible By 5")
	} else {
		fmt.Println("Number is not Divisible By 5")
	}
}
