package main

import "fmt"

func main() {
	s := "aBCdef"
	i := 0

	fmt.Println("Ganapati Bappa Morya.......")

	arr := []rune(s)
	//brr := make([]rune,30)

	for i = len(s)-1;i >= 0;i-- {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			fmt.Printf("%c",arr[i] - 32)
		} else if arr[i] >= 'A' && arr[i] <= 'Z'{ 
			fmt.Printf("%c",arr[i] + 32)
		} else {
			fmt.Printf("%c",arr[i])
		}
	}
}