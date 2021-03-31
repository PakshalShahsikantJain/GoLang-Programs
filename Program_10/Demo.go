package main

import "fmt"

func Print(Value, Value2 int) int {
	for i := 0; i < Value2; i++ {
		fmt.Print(Value, "\t")
	}

	return 0
}

func main() {

	var No1 int
	var No2 int
	fmt.Println("Ganapati Bappa Morya....")

	fmt.Println("Enter First Number To Print")
	fmt.Scan(&No1)

	fmt.Println("Enter Second Number For Iterations")
	fmt.Scan(&No2)

	Print(No1, No2)
}
