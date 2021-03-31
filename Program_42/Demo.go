package main

import "fmt"

func Display(No int) {
	i := 0

	if No < 0 {
		No = -No
	}
	for i = 0; i <= No; i++ {
		if i%2 == 1 {
			fmt.Print(i, "\t")
		}
	}
}
func main() {
	No := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
