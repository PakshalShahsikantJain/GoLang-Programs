package main

import "fmt"

func Display(No, No2 int) {
	i := 0

	if No2 < No {
		fmt.Println("Invalid Range")
	} else {
		for i = No2; i >= No; i-- {
			fmt.Print(i, " ")
		}
	}
}
func main() {
	No := 0
	No2 := 0

	fmt.Println("Enter Number To Start")
	fmt.Scan(&No)

	fmt.Println("Enter Number To End")
	fmt.Scan(&No2)

	Display(No, No2)
}
