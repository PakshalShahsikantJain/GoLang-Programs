package main

import "fmt"

func Print(Value int) int {
	for i := 0; i < Value; i++ {
		fmt.Print("* \t")
	}

	return 0
}
func main() {

	var No int
	fmt.Println("Enter Number of Iterations")
	fmt.Scan(&No)

	Print(No)
}
