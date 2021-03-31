package main

import "fmt"

func ODDMult(No int,arr[] int) int  {
	i := 0
	Mult := 1

	for i = 0;i < No;i++ {
		if arr[i] % 2 == 1 {
			Mult = Mult * arr[i]
		}
	}

	return Mult
}
func main(){
	No := 0
	i := 0
	ret := 0

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i = 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	ret = ODDMult(No,arr)
	fmt.Println("Multiplication of ODD Elements is :",ret)
}