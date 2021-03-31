package main

import "fmt"

func Display(ch rune) {
	i := '0'
	fmt.Println("Output :")
	if ch >= 'A' && ch <= 'Z' {
		for i = ch;i <='Z';i++ {
			fmt.Printf("%c  ",i)
		}
	} else if ch >= 'a' && ch <= 'z' {
		for i = ch;i >= 'a';i-- {
			fmt.Printf("%c  ",i)
		}
	} else {
		return
	}
}
func main() {
	ch := '0'

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Alpahabet")
	fmt.Scanf("%c",&ch)

	Display(ch)
}