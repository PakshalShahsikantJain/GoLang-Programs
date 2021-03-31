package main

import "fmt"

func Min(No int,arr[] int) int {
	i := 0
	iMin := 0

	iMin = arr[i]
	for i = 0;i < No;i++ {
		if arr[i] < iMin {
			iMin = arr[i]
		}
	}

	return iMin
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

	ret = Min(No,arr)
	fmt.Println("Smallest Element in Array is :",ret)
}