package main

import "fmt"

func Display(No, No2 int) {
	i := 0

	if No2 < No {
		fmt.Println("Invalid Number")
	}

	for i = No; i <= No2; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
func main() {
	fmt.Println("Jay Ganesh...")

	No := 0
	No2 := 0
	fmt.Println("Enter Number To Start")
	fmt.Scan(&No)

	fmt.Println("Enter Nummber To End")
	fmt.Scan(&No2)

	Display(No, No2)
}
