package main

import "fmt"

func Addition(No, No2 int) int {
	i := 0
	Sum := 0

	if No < 0 || No2 < 0 {
		fmt.Println("Invalid Range")
	} else if No2 < No {
		fmt.Println("Invalid Range")
	} else {
		for i = No; i <= No2; i++ {
			if i%2 == 0 {
				Sum = Sum + i
			}
		}
	}

	return Sum
}
func main() {
	fmt.Println("Jay Ganesh.....")

	No := 0
	No2 := 0
	ret := 0

	fmt.Println("Enter Number to Start")
	fmt.Scan(&No)

	fmt.Println("Enter Number To End")
	fmt.Scan(&No2)

	ret = Addition(No, No2)
	fmt.Println("Addition of Number's In Range is : ", ret)
}
