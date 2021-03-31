package main

import "fmt"

func main() {
	fmt.Println("Jay Ganesh.........")

	i := 0

	arr := make([]rune,30)
	brr := make([]rune,30)

	fmt.Println("Enter String")

	for arr[i] != '\n' {
		fmt.Scanf("%c",&arr[i])
		brr[i] = arr[i]
		fmt.Printf("%c",brr[i])
		i++
		i--
	}
}