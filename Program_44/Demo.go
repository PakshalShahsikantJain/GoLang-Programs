package main

import "fmt"

func Area(Radius float64) float64 {
	PI := 3.14
	Area := 0.0

	Area = PI * Radius * Radius

	return Area
}
func main() {
	fmt.Println("Jay Ganesh.......")

	No := 0.0
	ret := 0.0

	fmt.Println("Enter Radius of Circle")
	fmt.Scan(&No)

	ret = Area(No)
	fmt.Println("Area of Circle is : ", ret)
}
