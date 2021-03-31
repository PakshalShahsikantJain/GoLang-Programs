package main

import "fmt"

func Compare(s string,s2 string) {
	
	if s == s2 {
		fmt.Println("Strings are Same")
	} else {
		fmt.Println("Strings are Not Same")
	}
}
func main(){
	s := "Marvellous Infosystem"
	s2 := "Marvellous Infosystem"

	fmt.Println("Jay Ganesh......")

	Compare(s,s2)
}