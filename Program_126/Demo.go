package main

import "fmt"

func Check(ch rune) bool {
	if ch >= '!' && ch <= '/' {
		return true
	} else {
		return false
	}
}
func main() {
	ch := '0'
	bret := false 

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Alphabet")
	fmt.Scanf("%c",&ch)

	bret = Check(ch)

	if bret == true {
		fmt.Println("It is Special Character")
	} else {
		fmt.Println("It is Not a Special Character")
	}
}