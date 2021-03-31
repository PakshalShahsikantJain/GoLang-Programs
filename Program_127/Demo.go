package main

import "fmt"

func Display(ch rune) {
	fmt.Printf("%d %x %o %c",ch,ch,ch,ch)
}
func main() {
	ch := '0'

	fmt.Println("Jay Ganesh........")
	fmt.Println("Enter Alphabet To Remove Information")
	fmt.Scanf("%c",&ch)

	Display(ch)
}