package main

import "fmt"

func Display(s string) {
	i := 0

	arr := []rune(s)

	for i < len(s) {
		if arr[i] >= '0' && arr[i] <= '9' {
			fmt.Printf("%c",arr[i])
		}
		i++
	}
}
func main() {
	s := "marve89llous121"
	fmt.Println("Jay Ganesh.........")

	Display(s)
}