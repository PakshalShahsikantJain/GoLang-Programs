package main

import "fmt"

func Add(No int,arr[] int) {
	i := 0
	Sum := 0
	Digit := 0
	iNo := 0
	for i = 0;i < No;i++ {
			iNo = arr[i]
			Sum = 0
			for iNo != 0 {
			Digit = iNo % 10
			Sum = Sum + Digit
			iNo = iNo / 10
		}

		fmt.Print(Sum,"\t")
	}
}
func main() {
	No := 0
	i := 0

	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	Add(No,arr)
}