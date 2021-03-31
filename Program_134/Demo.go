package main

import "fmt"

func Convert(s string) {
	i := 0

	arr := []rune(s)

	for i < len(s) {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			fmt.Printf("%c",arr[i] - 32)
		} else {
			fmt.Printf("%c",arr[i])
		}
		i++
	}
}
func main() {
	s := "Marvellous Infosystem OS"
	fmt.Println("Jay Ganesh.......")

	Convert(s)
}