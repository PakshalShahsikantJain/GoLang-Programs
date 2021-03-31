package main

import "fmt"

func Check(No int,arr[] int,NO int) bool {
	i := 0

	for i = 0;i < No-1;i++ {
		if arr[i] == NO {
			break
		}
	}

	if arr[i] == NO {
		return true
	} else {
		return false
	}
}
func main() {
	No := 0
	i := 0
	NO := 0
	bret := false

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Enter Element To Search")
	fmt.Scan(&NO)

	bret = Check(No,arr,NO)

	if bret == true {
		fmt.Println("NO is Present in Array")
	} else {
		fmt.Println("NO is Not Present in Array")
	}
}