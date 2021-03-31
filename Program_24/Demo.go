package main

import "fmt"

func DisplayRev(No int) {
	Digit := 0
	for No != 0 {
		Digit = No % 10
		fmt.Println(Digit)
		No = No / 10
	}
}
func main() {
	No := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	DisplayRev(No)
}
