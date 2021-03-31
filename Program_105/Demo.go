package main

import "fmt"

func Check(No int,arr[] int) bool {
	i := 0

	for i = 0;i < No-1;i++ {
		if arr[i] == 11 {
			break
		}
	}

	if arr[i] == 11 {
		return true
	} else {
		return false
	}
}
func main() {
	No := 0
	i := 0
	bret := false

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")

	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	bret = Check(No,arr)

	if bret == true {
		fmt.Println("Eleven is Present in Array")
	} else {
		fmt.Println("Eleven is Absent in Array")
	}
}