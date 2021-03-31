package main

import "fmt"

func Frequency(No int,arr[] int) int {
	i := 0
	iCnt := 0

	for i = 0;i < No;i++ {
		if arr[i] % 2 == 0 {
			iCnt++
		}
	}

	return iCnt
}
func main() {
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	ret = Frequency(No,arr)
	fmt.Println("Frequency of Even Elements is :",ret)
}