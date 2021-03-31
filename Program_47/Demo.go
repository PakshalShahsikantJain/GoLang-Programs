package main

import "fmt"

func Conversion(Fahrenheit float64) float64 {
	Celcius := 0.0

	Celcius = (Fahrenheit - 32) * 5 / 9

	return Celcius
}
func main() {
	fmt.Println("Jay Ganesh....")

	Temperature := 0.0
	ret := 0.0

	fmt.Println("Enter Temerature in Faranheit")
	fmt.Scan(&Temperature)

	ret = Conversion(Temperature)
	fmt.Println("Conversion of Fahrenheit into Celcius is : ", ret)
}
