package main

import "fmt"

func DisplayPattern(No int) {
	i := 0

	for i = 1; i <= No ;i++ {
		fmt.Print(i,"\t","*\t")
	}
}
func main(){
	No := 0
	fmt.Println("Jay Ganesh....")

	fmt.Println("Enter Number")
	fmt.Scan(&No)

	DisplayPattern(No)
}