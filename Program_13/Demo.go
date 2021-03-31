package main

import "fmt"

func factors(No int) int {

	for iCnt := 1; iCnt <= No/2; iCnt++ {
		if iCnt%2 == 0 && No%iCnt == 0 {
			fmt.Println(iCnt)
		}
	}

	return 0
}
func main() {
	fmt.Println("Jay Ganesh....")

	No := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	factors(No)
}
