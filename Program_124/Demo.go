package main

import "fmt"

func Convert(ch rune) {
	if ch >= 'A' && ch <= 'Z'{
		fmt.Println("Converted Alpahabet is :")
		fmt.Printf("%c",ch+32)
	} else if ch >= 'a' && ch <= 'z'{
		fmt.Println("Converted Alpahabet is :")
		fmt.Printf("%c",ch-32)
	} else {
		fmt.Printf("%c",ch)
	}
}
func main() {
	ch := '0'

	fmt.Println("Jay Ganesh..........")
	fmt.Println("Enter Alpahabet To Convert....")
	fmt.Scanf("%c",&ch)

	Convert(ch)
}