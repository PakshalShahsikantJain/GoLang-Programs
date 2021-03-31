package main

import "fmt"

func Check(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true 
	} else {
		return false
	}
}
func main() {
	ch := '0'
	bret := false

	fmt.Println("Jay Ganesh.....")
	fmt.Println("Enter Alphabet To Check")
	fmt.Scanf("%c",&ch)

	bret = Check(ch)
	if bret == true {
		fmt.Println("Alphabet is in Small case")
	} else {
		fmt.Println("Alphabet is Not in Small case")
	}
}