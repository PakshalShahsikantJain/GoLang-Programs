package main

import "fmt"

func Display(No int) {
	i := 0

	fmt.Print("Output : ")
	for i = -No; i <= No; i++ {
		fmt.Print(i, "\t")
	}
}
func main() {
	fmt.Println("Jay Ganesh......")

	No := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
