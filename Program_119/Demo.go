package main

import "fmt"

func Check(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true 
	} else {
		return false
	}
}
func main() {
	ch := '0'
	bret := false

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Alphabet To Check....")
	fmt.Scanf("%c",&ch)

	bret = Check(ch)
	if bret == true {
		fmt.Println("The Alphabet is in UpperCase")
	} else {
		fmt.Println("The Alphabet May be in LowerCase or Entered Character May be Not Alphabet")
	}
}