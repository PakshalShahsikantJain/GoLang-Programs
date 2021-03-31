package main

import "fmt"

func Display(No int) {

	if No < 0{
		No = -No
	}
	for i:= 0; i < No ;i++{
		fmt.Print("*\t")
	}

	for j:= 0; j < No ;j++{
		fmt.Print("#\t")
	}

}
func main() {
	fmt.Println("Jay Ganesh.....")

	No := 0
	fmt.Println("Enter Number")
	fmt.Scan(&No)

	Display(No)
}
