package main

import "fmt"

func Display(No int) {

	if No < 50 {
		fmt.Println("Small")
	} else if No > 50 && No < 100 {
		fmt.Println("Medium")
	} else if No > 100 {
		fmt.Println("Large")
	}
}
func main() {
	No := 0

	fmt.Println("Jay Ganesh...")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
