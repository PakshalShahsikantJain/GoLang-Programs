package main

import "fmt"

func main() {
	s := "Pakshal Karande"
	i := 0

	fmt.Println("Ganapati Bappa Morya.......")

	arr := []rune(s)
	brr := make([]rune,30)

	for i = len(s)-1;i >= 0;i-- {
		brr[i] = arr[i]
		fmt.Printf("%c",brr[i])
	}
}