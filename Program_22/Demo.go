package main

import "fmt"

func Check(No, No2 int) bool {
	if No == No2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Jay Ganesh.......")
	No := 0
	No2 := 0
	bRet := false

	fmt.Println("Enter Number's To Check")
	fmt.Scanf("%d %d", &No, &No2)

	bRet = Check(No, No2)

	if bRet == true {
		fmt.Println("Number is Equal")
	} else {
		fmt.Println("Number is Not Equal")
	}
}
