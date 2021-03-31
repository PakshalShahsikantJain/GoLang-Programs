package main

import "fmt"

func Check(ch rune) bool {
	if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
		return true
	} else {
		return false
	}
}
func main() {
	ch := '0'
	bret := false

	fmt.Println("Jay Ganesh.......")
	fmt.Println("Enter Character To Check Whether it is Alpahabet or Not")
	fmt.Scanf("%c",&ch)

	bret = Check(ch)

	if bret == true {
		fmt.Println("It is Alpahabet")
	} else {
		fmt.Println("It is Not Alpahabet")
	}
}