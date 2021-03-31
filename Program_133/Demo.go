package main

import "fmt"

func Convert(s string) {
	i := 0

	arr := []rune(s)

	fmt.Println("Modified String is :")
	for i < len(s) {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			fmt.Printf("%c",arr[i] + 32)
		} else {
			fmt.Printf("%c",arr[i])
		}

		i++
	}
}
func main() {
	s := "Marvellous InfoSystem OS"
	fmt.Println("Jay Ganesh.........")
	fmt.Println(s)

	Convert(s)
}	