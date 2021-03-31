package main

import "fmt"

func Rectangle(Height, Width float64) float64 {
	Area := 0.0

	Area = Height * Width

	return Area
}
func main() {
	fmt.Println("Jay Ganesh...")

	No := 0.0
	No2 := 0.0

	ret := 0.0
	fmt.Println("Enter Height To Calculate")
	fmt.Scan(&No)

	fmt.Println("Enter Width To Calculate")
	fmt.Scan(&No2)

	ret = Rectangle(No, No2)
	fmt.Println("Area of Rectangle is : ", ret)
}
