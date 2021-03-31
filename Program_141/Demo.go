package main

import "fmt"

func Reverse(s string) {
	i := 0

	arr := []rune(s)

	for i = len(s) - 1; i >= 0;i-- {
		fmt.Printf("%c",arr[i])
	}  	
}
func main() {
	s := "abba"

	fmt.Println("Jay Ganesh...")

	Reverse(s)
}