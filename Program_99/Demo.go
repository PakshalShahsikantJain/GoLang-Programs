package main

import "fmt"

func Display(No int,arr[] int) {
	i := 0

	fmt.Println("Output :")
	for i = 0;i < No;i++ {
		if arr[i] % 5 == 0 {
			fmt.Print(arr[i],"\t")
		}
	}
}
func main() {
	No := 0

	fmt.Println("Jay Ganesh........")

	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i := 0;i < No;i++{
		fmt.Scan(&arr[i])
	}

	Display(No,arr)
}