package main

import "fmt"

func FirstOCC(No int,arr[] int,NO int) int {
	i := 0

	for i = 0;i < No-1;i++ {
		if arr[i] == NO {
			break
		}
	}

	if arr[i] == NO {
		return i
	} else { 
		return -1
	}
}
func main() {
	No := 0
	NO := 0
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

	fmt.Println("Enter Element To Search FirstOCC of :")
	fmt.Scan(&NO)

	ret = FirstOCC(No,arr,NO)

	if ret == -1 {
		fmt.Println("There is No Such Element in Array")		
	} else {
		fmt.Println("FirstOCC of Entered Element is At Index",ret)
	}

}	