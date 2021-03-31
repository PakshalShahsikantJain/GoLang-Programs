package main

import "fmt"

func Max(No int,arr[] int) int {
	i := 0
	iMax := 0

	iMax = arr[i]
	for i = 0;i < No;i++ {
		if arr[i] > iMax {
			iMax = arr[i]
		}
	}

	return iMax
}
func main(){
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh....")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	ret = Max(No,arr)
	fmt.Println("Max Element in Array is :",ret)
}