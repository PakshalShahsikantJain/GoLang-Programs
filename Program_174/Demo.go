package main

import "fmt"

func Pattern(s string) {
	i := 0
	j := 0

	arr := []rune(s)
	for i = 0;i < len(arr);i++ {
		fmt.Println()
		for j = 0;j < len(arr);j++ {
			if arr[j] >= 'A' && arr[j] <= 'Z' {
				fmt.Printf("%c\t",arr[j] + 32)
			} else {
				fmt.Printf("%c\t",arr[j])
			}
		}
	}
}
func main() {
	s := "Marvellous"
	fmt.Println("Jay Ganesh......")

	Pattern(s)
}