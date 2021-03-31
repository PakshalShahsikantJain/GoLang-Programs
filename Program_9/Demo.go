package main

import "fmt"

func Print(No int) int {
	if No < 10 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Demo")
	}

	return 0
}
func main() {

	var No int
	fmt.Println("Jay Ganesh....")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Print(No)
}
