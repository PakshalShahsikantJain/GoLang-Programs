package main

import "fmt"

func Pattern(No int) {
	i := 0

	if No < 0 {
		No = -No
	}
	for i = 1; i <= No; i++ {
		fmt.Print("$\t*\t")
	}
}
func main() {
	No := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Pattern(No)
}
