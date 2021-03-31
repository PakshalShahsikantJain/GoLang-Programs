package main

import "fmt"

func Check(No int) bool {
	Digit := 0

	for No != 0 {
		Digit = No % 10
		if Digit == 0 {
			break
		}
		No = No / 10
	}

	if Digit == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	No := 0
	bRet := false
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	bRet = Check(No)
	if bRet == true {
		fmt.Println("Zero is present")
	} else {
		fmt.Println("Zero is Not Present")
	}
}
