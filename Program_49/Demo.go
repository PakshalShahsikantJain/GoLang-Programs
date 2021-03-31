package main

import "fmt"

func Display(No, No2 int) {
	i := 0

	if No2 < No {
		fmt.Println("Invalid Range")
	}
	for i = No; i <= No2; i++ {
		fmt.Print(i, " ")
	}
}
func main() {
	fmt.Println("Ganapati Bappa Morya.......")

	No := 0
	No2 := 0
	fmt.Println("Enter First Number To Start")
	fmt.Scan(&No)

	fmt.Println("Enter Second Number To End")
	fmt.Scan(&No2)

	Display(No, No2)
}
