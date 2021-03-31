package main

import "fmt"

func Count(No int) int {
	Digit := 0
	iCnt := 0

	for No != 0 {
		Digit = No % 10
		if Digit == 4 {
			iCnt++
		}
		No = No / 10
	}

	return iCnt
}
func main() {
	fmt.Println("Jay Ganesh....")
	No := 0
	ret := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Count(No)
	fmt.Println("Count of No of 4 is : ", ret)
}
