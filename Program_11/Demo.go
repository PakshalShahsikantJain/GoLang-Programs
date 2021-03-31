package main

import "fmt"

func Check(Value int) bool {
	if Value%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	var No int
	fmt.Println("Ganapati Bappa Morya...")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	bRet := Check(No)
	if bRet == true {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}
}
