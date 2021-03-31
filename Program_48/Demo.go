package main

import "fmt"

func Conversion(SquareFeet float64) float64 {
	SquareMeter := 0.0929
	Conversion := 0.0

	Conversion = SquareFeet * SquareMeter

	return Conversion
}
func main() {
	fmt.Println("Jay Ganesh.....")

	No := 0.0
	ret := 0.0

	fmt.Println("Enter Square Feet")
	fmt.Scan(&No)

	ret = Conversion(No)
	fmt.Println("Conersion is : ", ret)
}
