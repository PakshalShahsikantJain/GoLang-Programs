package main

import "fmt"

func Display(No, No2 int) {
	i := 0
	Mult := 0

	if No < 0 || No2 < 0 {
		No = -No
		No2 = -No2
	}
	for i = 1; i <= No2; i++ {
		Mult = No * i
		fmt.Print(Mult, "\t")
	}
}
func main() {
	fmt.Println("Jay Ganesh...")

	No := 0
	No2 := 0

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	fmt.Println("Enter Number upto You Want To Multiply")
	fmt.Scan(&No2)

	Display(No, No2)
}
