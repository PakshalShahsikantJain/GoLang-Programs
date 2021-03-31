package main

import "fmt"

func Count(No int) int {
	iCnt := 0
	Digit := 0

	for No != 0 {
		Digit = No % 10
		if Digit == 2 {
			iCnt++
		}
		No = No / 10
	}

	return iCnt
}

func main() {
	fmt.Println("Jay Ganesh.....")
	ret := 0

	No := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Count(No)

	fmt.Println("Count of Number 2 is : ", ret)
}
