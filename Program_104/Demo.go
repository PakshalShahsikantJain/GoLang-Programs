package main

import "fmt"

func Diff(No int,arr[] int) int {
	i := 0
	iCntEven := 0
	iCntOdd := 0
	Diff := 0

	for i = 0;i < No;i++ {
		if arr[i] % 2 == 0{
			iCntEven++
		} else {
			iCntOdd++
		}
	}

	fmt.Println("Output :")
	fmt.Println(iCntEven,"-",iCntOdd)
	Diff = iCntEven - iCntOdd

	return Diff
}
func main() {
	No := 0
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

	ret = Diff(No,arr)
	fmt.Println("Difference Between Even Frequency and Odd Frequency is :",ret)
}