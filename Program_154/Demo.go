package main

import "fmt"

func Check(No int,No2 int,arr []int) bool {
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

	fmt.Println("Ganapati Bappa Morya......")
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter How Many Bits You Want To Check")
	fmt.Scan(&No2)

	arr := make([]int,No2)

	fmt.Println("Enter Bits")
	for i = 0;i < No2;i++ {
		fmt.Scan(&arr[i])
		//fmt.Println(arr[i])
	}

	bret = Check(No,No2,arr)

	if bret == true {
		fmt.Println("One or More or All Bits are On")
	} else {
		fmt.Println("One or More or All Bits are OFF")
	}
}