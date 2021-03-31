package main

import "fmt"

func Difference(No int,arr[] int) int {
	i := 0
	SumEven := 0
	SumOdd := 0
	Diff := 0

	for i = 0;i < No;i++ {
		if arr[i] % 2 == 0 {
			SumEven = SumEven + arr[i]
		} else {
			SumOdd = SumOdd + arr[i]
		}
	}

	fmt.Println("Output :")
	fmt.Println(SumEven,"-",SumOdd)
	Diff = SumEven - SumOdd

	return Diff
}
func main() {
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter Size of Arr")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0; i < No;i++ {
		fmt.Scan(&arr[i])
	}

	/*fmt.Println("Elements of Array Are :")
	for i = 0; i < No;i++{
		fmt.Print(arr[i],"\t")
	}*/

	ret = Difference(No,arr)
	fmt.Println("Difference Between Even And Odd Element is :",ret)
}