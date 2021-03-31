package main

import "fmt"

func Toggel(No,No2 int,arr[] int) int {
	i := 0
	iResult := 0
	iMask := 0X00000001

	for i = 0;i < No2;i++ {
		iMask = iMask << arr[i]
		iResult = No ^ iMask
	} 

	return iResult
} 
func main() {
	No := 0
	No2 := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter How Many Bits You Want To Toggel")
	fmt.Scan(&No2)

	arr := make([]int,No2)

	fmt.Println("Enter Bits You Want To Toggel")
	for i = 0;i < No2 ; i++ {
		fmt.Scan(&arr[i])
		//fmt.Println(arr[i])
	}

	ret = Toggel(No,No2,arr)

	fmt.Println(ret)
}