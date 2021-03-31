package main

import "fmt"

func Display(Value int) int {
	for i := 1; i <= Value; i++ {
		if (2*i)%2 == 0 {
			fmt.Print(2*i, "\t")
		}
	}

	return 0
}
func main() {

	var No int
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
