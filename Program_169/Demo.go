package main

import "fmt"

func main() {
	s := "MarraM"
	i := 0

	arr := []rune(s)
	//brr := make([]rune,30)

	fmt.Println("Jay Ganesh......")

	for i = len(s) - 1;i >= 0;i-- {
		fmt.Printf("%c",arr[i])
	}
}