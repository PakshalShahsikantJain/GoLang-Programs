package main

import "fmt"

func Toggel(s string) {
	i := 0

	arr := []rune(s)

	fmt.Println("Output :")
	for i < len(s) {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			fmt.Printf("%c",arr[i] + 32)
		} else {
			fmt.Printf("%c",arr[i] - 32) 
		}
		i++
	}
}
func main() {
	s := "Marvellous Multi OS"
	fmt.Println("Jay Ganesh.......")
	fmt.Println("Input :")
	fmt.Println(s)

	Toggel(s)
}