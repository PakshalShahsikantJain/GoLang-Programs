package main

import "fmt"

func Convert(Kilometer int) int {
	Meter := 1000

	if Kilometer < 0 {
		return -1
	}
	Conversion := 0
	Conversion = Kilometer * Meter // if 1 KM = 1000 M

	return Conversion
}
func main() {
	fmt.Println("Jay Ganesh.....")

	No := 0
	ret := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	ret = Convert(No)
	fmt.Println("Conersion is : ", ret)
}
