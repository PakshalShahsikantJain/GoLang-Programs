package main

import "fmt"

func Toggel(No,No2 int) int {
	iMask := 0X00000001
	iResult := 0

	iMask = iMask << No2
	iResult = No ^ iMask

	return iResult
}
func main() {
	No := 0
	No2 := 0

	iRet := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter Position You Want To Toggel")
	fmt.Scan(&No2)

	iRet = Toggel(No,No2)

	fmt.Println("Modified No is :",iRet)
}