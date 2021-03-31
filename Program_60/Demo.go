package main

import "fmt"

func DisplayPattern(No int) {
	i := 0

	for i = No; i > 0 ;i--{
		fmt.Printf("%d\t#\t",i)
	}
}
func main(){
	No := 0
	fmt.Println("Jay Ganesh..........")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	DisplayPattern(No)
}