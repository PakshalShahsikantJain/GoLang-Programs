package main

import "fmt"

func Check(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true 
	} else {
		return false
	}
}
func main() {
	ch := '0'
	bret := false

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Character To Check Whether it is Digit or Not....")
	fmt.Scanf("%c",&ch)

	bret = Check(ch)
	if bret == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}