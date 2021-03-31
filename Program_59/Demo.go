package main

import "fmt"

func DisplayChar(No int) {
	var ch byte = 'A'

	for i:= 0; i < No; i++ {
		fmt.Printf("%c\t",ch)
		ch++
	}
}
func main(){
	No := 0

	fmt.Println("Jay Ganesh.......")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	DisplayChar(No)
}