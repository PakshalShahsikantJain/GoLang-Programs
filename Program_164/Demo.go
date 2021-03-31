package main

import "fmt"

func Check(No,No2 int,arr[] int) bool {
	i := 0
	iMask := 0X00000001
	iResult := 0

	for i = 0;i < No2;i++ {
		iMask = iMask << arr[i]
		iResult = No & iMask
		if iResult == iMask {
			break
		}
	}

	if iResult == iMask {
		return true 
	} else {
		return false
	}
}
func main() {
	No := 0
	No2 := 0 
	i := 0
	bret := false

	fmt.Println("Jay Ganesh.....")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter How Many Position You Want To Check")
	fmt.Scan(&No2)

	arr := make([]int,No2)

	fmt.Println("Enter Elements in Array")

	for i = 0;i < No2;i++ {
		fmt.Scan(&arr[i])
	}

	bret = Check(No,No2,arr)
	if bret == true {
		fmt.Println("Bits are on")
	} else {
		fmt.Println("Bits are Off")
	}
}