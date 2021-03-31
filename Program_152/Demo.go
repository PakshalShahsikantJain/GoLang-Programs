package main

import "fmt"

func Check(No,No2 int) bool {
	iMask := 0X00000001
	iResult := 0

	if No < 0 {
		No = -No
	}

	if (No2 < 1) || (No2 > 32) {
		return false
	}

	iMask = iMask << No2 
	iResult = No & iMask

	if iResult == iMask {
		return true
	} else {
		return false
	}
}

func main() {
	No := 0
	No2 := 0
	bret := false

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter Bit To Check Whether it is On or OFF")
	fmt.Scan(&No2)

	bret = Check(No,No2)

	if bret == true {
		fmt.Println("Bit is On")
	} else {
		fmt.Println("Bit is OFF")
	}
}