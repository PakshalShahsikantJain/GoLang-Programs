package main

import "fmt"

func Difference(No int,arr[] int) int {
	i := 0
	iMax := 0
	iMin := 0

	iMax = arr[i]
	iMin = arr[i]
	for i = 0;i < No;i++ {
		if arr[i] > iMax {
			iMax = arr[i]
		} else if arr[i] < iMin {
			iMin = arr[i]
		}
	}

	return iMax - iMin
}
func main() {
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh.....")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	ret = Difference(No,arr)
	fmt.Println("Difference Between Largest and Smallest Element is :",ret)
}