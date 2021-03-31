package main

import "fmt"

func Count(No int) int {
	iCnt := 0
	Value := 0

	if No < 0 {
		No = -No
	}

	for No != 0 {
		Value = No % 10
		if Value > 3 && Value < 7 {
			iCnt++
		}
		No = No / 10
	}

	return iCnt
}
func main() {
	fmt.Println("Jay Ganesh")

	No := 0
	ret := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Count(No)
	fmt.Println("Count of Numbers Between Three and seven is : ", ret)
}
