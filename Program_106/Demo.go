package main

import "fmt"

func Count(No int,arr[] int) int {
	iCnt := 0
	i := 0

	for i = 0;i < No;i++ {
		if arr[i] == 11 {
			iCnt++
		}
	}

	return iCnt
}
func main(){
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	ret = Count(No,arr)
	fmt.Println("Frequency of Eleven is :",ret)
}