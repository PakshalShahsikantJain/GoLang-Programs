package main

import "fmt"

func Count(No int) int {
	Value := 0
	iCnt := 0
	if No < 0 {
		No = -No
	}
	for No != 0 {
		Value = No % 10
		if Value%2 == 1 {
			iCnt++
		}
		No = No / 10
	}

	return iCnt
}
func main() {
	fmt.Println("Jay Ganesh...")
	No := 0
	ret := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Count(No)
	fmt.Println("Even number of Count is : ", ret)
}
