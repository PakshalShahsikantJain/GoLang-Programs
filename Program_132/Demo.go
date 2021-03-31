package main

import "fmt"

func Reverse(s string) {
	i := len(s) - 1

	arr := []rune(s)

	for len(s)  != 0 && i >= 0{
		fmt.Printf("%c",arr[i])
		i--
	}
}
func main() {
	s := " "
	fmt.Println("Jay Ganesh........")

	fmt.Println("Enter String")
	fmt.Scan(&s)
	
	Reverse(s)
}