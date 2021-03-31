package main

import "fmt"

func Display(No int,arr[] int,NO int,NO2 int) {
	i := 0

	fmt.Println("Output :")
	for i = 0;i < No;i++ {
		if arr[i] > NO && arr[i] < NO2{
			fmt.Print(arr[i],"\t")
		}
	}
}
func main(){
	No := 0
	NO := 0
	NO2 := 0

	fmt.Println("Jay Ganesh.........")
	fmt.Println("Enter Size of Array")
	fmt.Scan(&No)

	arr := make([]int,No)

	fmt.Println("Enter Elements in Array")
	for i := 0;i < No;i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Enter Element To Start")
	fmt.Scan(&NO)

	fmt.Println("Enter Element To End")
	fmt.Scan(&NO2)

	Display(No,arr,NO,NO2)
}