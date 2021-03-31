package main

import "fmt"

func Display(No int) {

	if No < 0 {
		No = -No
	}

	if No >= 0 && No <= 9 {
		switch No {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		}
	} else {
		fmt.Println("Invalid Number")
	}
}
func main() {
	fmt.Println("Jay Ganesh...")
	No := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
