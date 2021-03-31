package main

import "fmt"

func Factorial(No int) {
	Mult := 1
	i := 1

	if No < 0 {
		No = -No
	}

	for i = No; i > 0; i-- {
		Mult = Mult * i
	}

	fmt.Println(Mult)
}

func main() {
	fmt.Println("Jay Ganesh....")

	No := 0
	fmt.Println("Enter Number..")
	fmt.Scan(&No)

	Factorial(No)
}
