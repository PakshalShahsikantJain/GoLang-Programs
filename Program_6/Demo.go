package main

import "fmt"

func Print(No int) int {
	for i := No; i > 0; i-- {
		fmt.Print(i,"\t")
	}

	return 0
}
func main() {

	var No int
	fmt.Println("Enter Number For Iterations")
	fmt.Scan(&No)

	Print(No)

}
