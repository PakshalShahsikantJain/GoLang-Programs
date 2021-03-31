package main

import "fmt"

func Count(No int,arr[] int) {
	i := 0
	iCnt := 0
	No2 := 0

	fmt.Println("Output :")
	for i = 0;i < No;i++ {
		No2 = arr[i]

		iCnt = 0
		for No2 != 0 {
			No2 = No2 / 10
			iCnt++
		}

		if iCnt == 3 {
			fmt.Print(arr[i],"\t")
		}
	}
} 
func main() {
	No := 0
	i := 0
	fmt.Println("Jay Ganesh.......")

	fmt.Println("Enter Size of array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No ;i++ {
		fmt.Scan(&arr[i])
	}

	Count(No,arr)
}