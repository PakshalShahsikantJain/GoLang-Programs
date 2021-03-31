package main

import "fmt"

func Convert(s string) {
	i := 0

	arr := []rune(s)
	brr := make([]rune,30)

	fmt.Println("Output :")
	for i = 0;i < len(s);i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			brr[i] = arr[i]
			fmt.Printf("%c",brr[i] + 32)
		} else {
			brr[i] = arr[i]
			fmt.Printf("%c",brr[i])
		}
	}
}
func main() {
	s := "Marvellous Python 2"

	fmt.Println("Jay Ganesh............")
	fmt.Println("Input :")
	fmt.Println(s)

	Convert(s)
}