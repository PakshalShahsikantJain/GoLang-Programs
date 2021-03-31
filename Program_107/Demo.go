package main

import "fmt"

func Count(No int,arr[] int,NO int) int {
	i := 0
	iCnt := 0

	for i = 0;i < No;i++ {
		if arr[i] == NO {
			iCnt++
		}
	}

	return iCnt
}
func main() {
	No := 0
	NO := 0
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

	fmt.Println("Enter Element To Search in Array")
	fmt.Scan(&NO)

	ret = Count(No,arr,NO)
	fmt.Println("Frequecy of Entered Number is :",ret)
}