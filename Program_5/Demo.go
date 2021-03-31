package main

import "fmt"

func Print(No int) int {
	fmt.Println("Output : ")
	for i := 0; i < No; i++ {
		fmt.Println("Marvellous")
	}
	return 0
}

func main() {

	fmt.Println("Enter Number for Iteration")
	var No int
	fmt.Scan(&No)

	Print(No)
}
