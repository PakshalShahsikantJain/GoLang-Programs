package main 

import "fmt"

func Standard(No int) {
	switch No {
	case 1 :
		fmt.Println("Your Fees For 1st Standard is 8900")
	case 2 :
		fmt.Println("Your Fess For 2nd Standard is 12790")
	case 3 :
		fmt.Println("Your Fees For 3rd Standard is 21000")
	case 4 :
		fmt.Println("Your Fees For 4th Standard is 23450")
	case 5 :
		fmt.Println("Wrong Input Please Re-Enter Your Standard")
	}
}
func main() {
	No := 0

	fmt.Println("Jay Ganesh......")

	fmt.Println("Enter Your Standard")
	fmt.Scan(&No)

	Standard(No)
}